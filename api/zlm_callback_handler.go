package api

import (
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"ivmp-service/entity"
	"ivmp-service/service"
	"log"
	"net/http"
)

func InitCallbackRouter(r chi.Router) {
	r.Post("/index/hook/on_play", OnPlayHandler)
	r.Post("/index/hook/on_publish", OnPublishHandler)
	r.Post("/index/hook/on_record_mp4", OnRecordMp4Handler)
	r.Post("/index/hook/on_rtsp_realm", OnRtspRealmHandler)
	r.Post("/index/hook/on_rtsp_auth", OnRtspAuthHandler)
	r.Post("/index/hook/on_shell_login", OnShellLoginHandler)
	r.Post("/index/hook/on_stream_changed", OnStreamChangedHandler)
	r.Post("/index/hook/on_stream_none_reader", OnStreamNoneReaderHandler)
	r.Post("/index/hook/on_stream_not_found", OnStreamNotFoundHandler)
	r.Post("/index/hook/on_server_started", OnServerStartedHandler)
	r.Post("/index/hook/on_server_keepalive", OnServerKeepaliveHandler)
}

func OnPlayHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("OnPlay")
	//TODO
	HttpResult(w, entity.ZlmResp{
		Code: 0,
		Msg:  "success",
	})
}

func OnPublishHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("OnPublish")
	//TODO
	HttpResult(w, entity.OnPublishResp{
		Code:      0,
		Msg:       "success",
		EnableHls: false,
		EnableMp4: false,
	})
}

func OnRecordMp4Handler(w http.ResponseWriter, r *http.Request) {
	log.Println("OnRecordMp4")
	//TODO
	HttpResult(w, entity.ZlmResp{
		Code: 0,
		Msg:  "success",
	})
}

func OnRtspRealmHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("OnRtspRealm")
	//TODO
	HttpResult(w, entity.OnRtspRealmResp{
		Code:  0,
		Realm: "",
		Msg:   "success",
	})
}

func OnRtspAuthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("OnRtspAuth")
	//TODO
	HttpResult(w, entity.OnRtspAuthResp{
		Code:      0,
		Msg:       "success",
		Passwd:    "",
		Encrypted: false,
	})
}

func OnShellLoginHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("OnShellLogin")
	//TODO
	HttpResult(w, entity.ZlmResp{
		Code: 0,
		Msg:  "success",
	})
}

func OnStreamChangedHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("OnStreamChanged")
	//TODO
	HttpResult(w, entity.ZlmResp{
		Code: 0,
		Msg:  "success",
	})
}

func OnStreamNoneReaderHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("OnStreamNoneReader")
	//TODO 判断是否是AI的流？
	HttpResult(w, entity.OnStreamNoneReaderResp{
		Code:  0,
		Close: true,
	})
}

func OnStreamNotFoundHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("OnStreamNotFound")
	var req entity.OnStreamNotFoundReq
	err := common.ReadRequestBody(r, &req)
	resp := entity.ZlmResp{}
	if err != nil {
		resp.Code = 1
		resp.Msg = err.Error()
		HttpResult(w, resp)
		return
	}
	err = service.ProcessStreamNotFound(r.Context(), req.Stream)
	if err != nil {
		resp.Code = 2
		resp.Msg = err.Error()
		HttpResult(w, resp)
		return
	}

	HttpResult(w, entity.ZlmResp{
		Code: 0,
		Msg:  "success",
	})
}

func OnServerStartedHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("OnServerStart")
	//TODO
	HttpResult(w, entity.ZlmResp{
		Code: 0,
		Msg:  "success",
	})
}
func OnServerKeepaliveHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("OnServerKeepalive")
	//TODO
	HttpResult(w, entity.ZlmResp{
		Code: 0,
		Msg:  "success",
	})
}
