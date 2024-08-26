package service

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"github.com/dapr-platform/common"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"ivmp-service/config"
	"ivmp-service/entity"
	"ivmp-service/eventpub"
	"ivmp-service/model"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

func init() {
	common.RegisterUpsertBeforeHook("Camera", ProcessUpsertCamera)
}
func ProcessUpsertCamera(r *http.Request, in any) (out any, err error) {
	sub, _ := common.ExtractUserSub(r)

	inObj := in.(model.Camera)
	if inObj.CreatedBy == "" {
		inObj.CreatedBy = sub
	}

	if inObj.ID == "" {
		inObj.ID = common.NanoId()
		inObj.CreatedTime = common.LocalTime(time.Now())
	}
	inObj.UpdatedBy = sub
	inObj.UpdatedTime = common.LocalTime(time.Now())

	out = inObj
	return
}
func ImportCameras(ctx context.Context, sub, filePath string) (err error) {
	file, err := os.Open(filePath)
	if err != nil {
		common.Logger.Error(err.Error())
		return
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)
	cameras := make([]model.Camera, 0)
	rows, err := reader.ReadAll()
	if err != nil {
		common.Logger.Error(err.Error())
		err = errors.Wrap(err, "read all csv file error")
		return err
	}
	for i := 0; i < len(rows); i++ {
		row := rows[i]

		if i == 0 { //标题行
			continue
		}
		if len(row) == 0 {
			continue
		}
		identifier := row[0]
		camera, err := getCameraByIdentifier(ctx, identifier)
		if err != nil {
			common.Logger.Error(err)
			err = errors.Wrap(err, "get device error")
			return err
		}
		if camera == nil {
			camera = &model.Camera{
				ID:               common.GetMD5Hash(identifier),
				Identifier:       identifier,
				Name:             strings.TrimSpace(row[1]),
				IP:               strings.TrimSpace(row[2]),
				StreamType:       strings.TrimSpace(row[3]),
				StreamPort:       int32(cast.ToInt(row[4])),
				Username:         strings.TrimSpace(row[5]),
				Password:         strings.TrimSpace(row[6]),
				StreamPath:       strings.TrimSpace(row[7]),
				SecondStreamPath: strings.TrimSpace(row[8]),
				CreatedBy:        sub,
				UpdatedBy:        sub,
				CreatedTime:      common.LocalTime(time.Now()),
				UpdatedTime:      common.LocalTime(time.Now()),
			}
		} else {
			camera.UpdatedBy = sub
			camera.Name = strings.TrimSpace(row[1])
			camera.IP = strings.TrimSpace(row[2])
			camera.StreamType = strings.TrimSpace(row[3])
			camera.StreamPort = int32(cast.ToInt(row[4]))
			camera.Username = strings.TrimSpace(row[5])
			camera.Password = strings.TrimSpace(row[6])
			camera.StreamPath = strings.TrimSpace(row[7])
			camera.SecondStreamPath = strings.TrimSpace(row[8])
			camera.UpdatedTime = common.LocalTime(time.Now())
		}
		cameras = append(cameras, *camera)
	}
	err = common.DbBatchUpsert[model.Camera](ctx, common.GetDaprClient(), cameras, model.CameraTableInfo.Name, model.Camera_FIELD_NAME_id)
	if err != nil {
		common.Logger.Error(err.Error())
		err = errors.Wrap(err, "upsert device error")
		return err
	}

	return nil

}
func getCameraByIdentifier(ctx context.Context, identifier string) (*model.Camera, error) {
	qstr := "identifier=" + identifier
	result, err := common.DbQuery[model.Camera](ctx, common.GetDaprClient(), model.CameraTableInfo.Name, qstr)
	if err != nil {
		return nil, errors.WithMessage(err, "db query failed")
	}
	if len(result) == 0 {
		log.Println("identifier:", identifier, "not found")
		return nil, nil
	}
	if len(result) > 1 {
		log.Println("identifier:", identifier, "found more than one")
		eventpub.ConstructAndSendEvent(ctx, common.EventTypePlatform, identifier+" found more than one", identifier+" found more than one", common.EventStatusActive, common.EventLevelMinor, time.Now(), "ivmp-service", "ivmp-service", "ivmp-service")
	}
	return &result[0], nil
}

func checkCameraFieldValidate(ctx context.Context, camera *model.Camera) bool {
	if camera.IP == "" || camera.StreamPort == 0 || camera.Username == "" || camera.Password == "" || camera.StreamPath == "" {
		cstr, _ := json.Marshal(camera)
		eventpub.ConstructAndSendEvent(ctx, common.EventTypePlatform, camera.Name+"rtsp流信息不全", string(cstr), common.EventStatusActive, common.EventLevelMinor, time.Now(), "zlm-manager", "zlm-manager", "zlm-manager")
		return false
	}
	return true
}

func ProcessStreamNotFound(ctx context.Context, stream string) error {
	return processCamera(ctx, stream)
}
func processCamera(ctx context.Context, stream string) error {
	camera, err := getCameraByIdentifier(ctx, stream)
	if err != nil {
		eventpub.ConstructAndSendEvent(ctx, common.EventTypePlatform, stream+" get camera error", err.Error(), common.EventStatusActive, common.EventLevelMinor, time.Now(), "zlm-manager", "zlm-manager", "zlm-manager")
	}
	rtspUrl := ""
	if camera != nil {
		if checkCameraFieldValidate(ctx, camera) {
			pathSep := ""
			if !strings.HasPrefix(camera.StreamPath, "/") {
				pathSep = "/"
			}
			pwd := url.QueryEscape(camera.Password)
			pwd = camera.Password
			rtspUrl = "rtsp://" + camera.Username + ":" + pwd + "@" + camera.IP + ":" + strconv.Itoa(int(camera.StreamPort)) + pathSep + camera.StreamPath
			log.Println("stream:", stream, "rtspUrl:", rtspUrl)
		} else {
			log.Println("camera checkCameraFieldValidate false ")
			return errors.New("camera checkCameraFieldValidate false " + stream)
		}

	} else {
		log.Println("can't find camera by " + stream)
		return errors.New("can't find camera " + stream)
	}

	rtspUrl = url.QueryEscape(rtspUrl)
	log.Println("stream:", stream, "rtspUrl:", rtspUrl)
	return addStreamProxy(ctx, rtspUrl, stream)
}

func addStreamProxy(ctx context.Context, rtspUrl, stream string) error {
	method := "/index/api/addStreamProxy?secret=" + config.ZLM_SECRET + "&vhost=__default_vhost__&app=live&stream=" + stream + "&url=" + rtspUrl
	common.Logger.Debug("addStreamProxy method:", method)
	ret, err := common.GetDaprClient().InvokeMethod(ctx, config.ZLM_SERVICE_NAME, method, "GET")
	if err != nil {
		return errors.WithMessage(err, "dapr invoke method failed")
	}
	var resp entity.ZlmResp
	err = json.Unmarshal(ret, &resp)
	if err != nil {
		return errors.WithMessage(err, "unmarshal resp failed")
	}
	if resp.Code != 0 {
		return errors.New(resp.Msg)
	}
	return nil
}
