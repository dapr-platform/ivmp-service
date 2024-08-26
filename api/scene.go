package api

import (
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"ivmp-service/model"
	"net/http"
	"strings"
)

func InitSceneRoute(r chi.Router) {
	r.Get(common.BASE_CONTEXT+"/scene/page", ScenePageListHandler)
	r.Get(common.BASE_CONTEXT+"/scene", SceneListHandler)
	r.Post(common.BASE_CONTEXT+"/scene", UpsertSceneHandler)
	r.Delete(common.BASE_CONTEXT+"/scene/{id}", DeleteSceneHandler)
	r.Post(common.BASE_CONTEXT+"/scene/batch-delete", batchDeleteSceneHandler)
	r.Post(common.BASE_CONTEXT+"/scene/batch-upsert", batchUpsertSceneHandler)
	r.Get(common.BASE_CONTEXT+"/scene/groupby", SceneGroupbyHandler)
}

// @Summary GroupBy
// @Description GroupBy, for example,  _select=level, then return  {level_val1:sum1,level_val2:sum2}, _where can input status=0
// @Tags Scene
// @Param _select query string true "_select"
// @Param _where query string false "_where"
// @Produce  json
// @Success 200 {object} common.Response{data=[]map[string]any} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /scene/groupby [get]
func SceneGroupbyHandler(w http.ResponseWriter, r *http.Request) {

	common.CommonGroupby(w, r, common.GetDaprClient(), "o_scene")
}

// @Summary batch update
// @Description batch update
// @Tags Scene
// @Accept  json
// @Param entities body []map[string]any true "objects array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /scene/batch-upsert [post]
func batchUpsertSceneHandler(w http.ResponseWriter, r *http.Request) {

	var entities []map[string]any
	err := common.ReadRequestBody(r, &entities)
	if err != nil {
		common.HttpResult(w, common.ErrParam)
		return
	}
	if len(entities) == 0 {
		common.HttpResult(w, common.ErrParam)
		return
	}

	err = common.DbBatchUpsert[map[string]any](r.Context(), common.GetDaprClient(), entities, model.SceneTableInfo.Name, model.Scene_FIELD_NAME_id)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags Scene
// @Param _page query int true "current page"
// @Param _page_size query int true "page size"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param name query string false "name"
// @Param created_by query string false "created_by"
// @Param created_time query string false "created_time"
// @Param updated_by query string false "updated_by"
// @Param updated_time query string false "updated_time"
// @Param type query string false "type"
// @Param config query string false "config"
// @Param description query string false "description"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]model.Scene}} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /scene/page [get]
func ScenePageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam)
		return
	}
	common.CommonPageQuery[model.Scene](w, r, common.GetDaprClient(), "o_scene", "id")

}

// @Summary query objects
// @Description query objects
// @Tags Scene
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param name query string false "name"
// @Param created_by query string false "created_by"
// @Param created_time query string false "created_time"
// @Param updated_by query string false "updated_by"
// @Param updated_time query string false "updated_time"
// @Param type query string false "type"
// @Param config query string false "config"
// @Param description query string false "description"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.Scene} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /scene [get]
func SceneListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.Scene](w, r, common.GetDaprClient(), "o_scene", "id")
}

// @Summary save
// @Description save
// @Tags Scene
// @Accept       json
// @Param item body model.Scene true "object"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Scene} "object"
// @Failure 500 {object} common.Response ""
// @Router /scene [post]
func UpsertSceneHandler(w http.ResponseWriter, r *http.Request) {
	var val model.Scene
	err := common.ReadRequestBody(r, &val)
	if err != nil {
		common.HttpResult(w, common.ErrParam)
		return
	}
	beforeHook, exists := common.GetUpsertBeforeHook("Scene")
	if exists {
		v, err1 := beforeHook(r, val)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
		val = v.(model.Scene)
	}

	err = common.DbUpsert[model.Scene](r.Context(), common.GetDaprClient(), val, model.SceneTableInfo.Name, "id")
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK.WithData(val))
}

// @Summary delete
// @Description delete
// @Tags Scene
// @Param id  path string true "实例id"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Scene} "object"
// @Failure 500 {object} common.Response ""
// @Router /scene/{id} [delete]
func DeleteSceneHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	beforeHook, exists := common.GetDeleteBeforeHook("Scene")
	if exists {
		_, err1 := beforeHook(r, id)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	common.CommonDelete(w, r, common.GetDaprClient(), "o_scene", "id", "id")
}

// @Summary batch delete
// @Description batch delete
// @Tags Scene
// @Accept  json
// @Param ids body []string true "id array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /scene/batch-delete [post]
func batchDeleteSceneHandler(w http.ResponseWriter, r *http.Request) {

	var ids []string
	err := common.ReadRequestBody(r, &ids)
	if err != nil {
		common.HttpResult(w, common.ErrParam)
		return
	}
	if len(ids) == 0 {
		common.HttpResult(w, common.ErrParam)
		return
	}
	beforeHook, exists := common.GetBatchDeleteBeforeHook("Scene")
	if exists {
		_, err1 := beforeHook(r, ids)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	idstr := strings.Join(ids, ",")
	err = common.DbDeleteByOps(r.Context(), common.GetDaprClient(), "o_scene", []string{"id"}, []string{"in"}, []any{idstr})
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}
