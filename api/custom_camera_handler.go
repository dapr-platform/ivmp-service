package api

import (
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"io"
	"ivmp-service/service"
	"net/http"
	"os"
)

func InitCustomCameraRoute(r chi.Router) {
	r.Post(common.BASE_CONTEXT+"/camera/import-cameras", importExcelHandler)

}

// @Summary 导入摄像头
// @Description 导入摄像头
// @Tags Camera
// @Param csv_data body string true "csv文件字符串"
// @Produce  application/json
// @Success 200 {object} common.Response "{"status":0,"data":{},"msg":"success"}"
// @Failure 500 {object} common.Response "错误code和错误信息"
// @Router /camera/import-cameras [post]
func importExcelHandler(w http.ResponseWriter, r *http.Request) {
	sub, _ := common.ExtractUserSub(r)

	newName := common.NanoId() + ".csv"
	common.Logger.Debug("newName=", newName)
	f, err := os.OpenFile(newName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		common.Logger.Error("打开文件失败", err)
		common.HttpResult(w, common.ErrService.AppendMsg("open file error"))
		return
	}

	io.Copy(f, r.Body)
	err = service.ImportCameras(r.Context(), sub, newName)
	defer func() {
		r.Body.Close()
		f.Close()
		os.Remove(newName)
	}()
	if err != nil {
		common.Logger.Error("处理文件失败", err)
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	} else {
		//event.ConstructAndSendEvent(context.Background(), common.EventTypePlatform, common.EventSubTypeService, "process excel "+handler.Filename+" error", "", common.EventStatusClosed, common.EventLevelMajor, time.Now(), "ops-service", "ops-service", "db-excel-upload")

	}

	common.HttpResult(w, common.OK)
}
