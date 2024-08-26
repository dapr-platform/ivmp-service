package api

import "github.com/go-chi/chi/v5"

func InitRoute(r chi.Router) {
	InitCameraRoute(r)
	InitSceneRoute(r)
	InitCallbackRouter(r)
	InitCustomCameraRoute(r)
	InitAi_modelRoute(r)
}
