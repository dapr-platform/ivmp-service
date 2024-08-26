package api

import (
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"ivmp-service/model"
	"net/http"
	"strings"
)

func InitAi_modelRoute(r chi.Router) {
	r.Get(common.BASE_CONTEXT+"/ai-model/page", Ai_modelPageListHandler)
	r.Get(common.BASE_CONTEXT+"/ai-model", Ai_modelListHandler)
	r.Post(common.BASE_CONTEXT+"/ai-model", UpsertAi_modelHandler)
	r.Delete(common.BASE_CONTEXT+"/ai-model/{id}", DeleteAi_modelHandler)
	r.Post(common.BASE_CONTEXT+"/ai-model/batch-delete", batchDeleteAi_modelHandler)
	r.Post(common.BASE_CONTEXT+"/ai-model/batch-upsert", batchUpsertAi_modelHandler)
	r.Get(common.BASE_CONTEXT+"/ai-model/groupby", Ai_modelGroupbyHandler)
}

// @Summary GroupBy
// @Description GroupBy, for example,  _select=level, then return  {level_val1:sum1,level_val2:sum2}, _where can input status=0
// @Tags Ai_model
// @Param _select query string true "_select"
// @Param _where query string false "_where"
// @Produce  json
// @Success 200 {object} common.Response{data=[]map[string]any} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /ai-model/groupby [get]
func Ai_modelGroupbyHandler(w http.ResponseWriter, r *http.Request) {

	common.CommonGroupby(w, r, common.GetDaprClient(), "o_ai_model")
}

// @Summary batch update
// @Description batch update
// @Tags Ai_model
// @Accept  json
// @Param entities body []map[string]any true "objects array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /ai-model/batch-upsert [post]
func batchUpsertAi_modelHandler(w http.ResponseWriter, r *http.Request) {

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

	err = common.DbBatchUpsert[map[string]any](r.Context(), common.GetDaprClient(), entities, model.Ai_modelTableInfo.Name, model.Ai_model_FIELD_NAME_id)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags Ai_model
// @Param _page query int true "current page"
// @Param _page_size query int true "page size"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param name query string false "name"
// @Param created_by query string false "created_by"
// @Param created_time query string false "created_time"
// @Param updated_by query string false "updated_by"
// @Param updated_time query string false "updated_time"
// @Param config query string false "config"
// @Param description query string false "description"
// @Param version query string false "version"
// @Param file_ext query string false "file_ext"
// @Param type query string false "type"
// @Param sub_type query string false "sub_type"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]model.Ai_model}} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /ai-model/page [get]
func Ai_modelPageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam)
		return
	}
	common.CommonPageQuery[model.Ai_model](w, r, common.GetDaprClient(), "o_ai_model", "id")

}

// @Summary query objects
// @Description query objects
// @Tags Ai_model
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param name query string false "name"
// @Param created_by query string false "created_by"
// @Param created_time query string false "created_time"
// @Param updated_by query string false "updated_by"
// @Param updated_time query string false "updated_time"
// @Param config query string false "config"
// @Param description query string false "description"
// @Param version query string false "version"
// @Param file_ext query string false "file_ext"
// @Param type query string false "type"
// @Param sub_type query string false "sub_type"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.Ai_model} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /ai-model [get]
func Ai_modelListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.Ai_model](w, r, common.GetDaprClient(), "o_ai_model", "id")
}

// @Summary save
// @Description save
// @Tags Ai_model
// @Accept       json
// @Param item body model.Ai_model true "object"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Ai_model} "object"
// @Failure 500 {object} common.Response ""
// @Router /ai-model [post]
func UpsertAi_modelHandler(w http.ResponseWriter, r *http.Request) {
	var val model.Ai_model
	err := common.ReadRequestBody(r, &val)
	if err != nil {
		common.HttpResult(w, common.ErrParam)
		return
	}
	beforeHook, exists := common.GetUpsertBeforeHook("Ai_model")
	if exists {
		v, err1 := beforeHook(r, val)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
		val = v.(model.Ai_model)
	}

	err = common.DbUpsert[model.Ai_model](r.Context(), common.GetDaprClient(), val, model.Ai_modelTableInfo.Name, "id")
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK.WithData(val))
}

// @Summary delete
// @Description delete
// @Tags Ai_model
// @Param id  path string true "实例id"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Ai_model} "object"
// @Failure 500 {object} common.Response ""
// @Router /ai-model/{id} [delete]
func DeleteAi_modelHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	beforeHook, exists := common.GetDeleteBeforeHook("Ai_model")
	if exists {
		_, err1 := beforeHook(r, id)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	common.CommonDelete(w, r, common.GetDaprClient(), "o_ai_model", "id", "id")
}

// @Summary batch delete
// @Description batch delete
// @Tags Ai_model
// @Accept  json
// @Param ids body []string true "id array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /ai-model/batch-delete [post]
func batchDeleteAi_modelHandler(w http.ResponseWriter, r *http.Request) {

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
	beforeHook, exists := common.GetBatchDeleteBeforeHook("Ai_model")
	if exists {
		_, err1 := beforeHook(r, ids)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	idstr := strings.Join(ids, ",")
	err = common.DbDeleteByOps(r.Context(), common.GetDaprClient(), "o_ai_model", []string{"id"}, []string{"in"}, []any{idstr})
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}
