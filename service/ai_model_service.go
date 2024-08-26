package service

import (
	"github.com/dapr-platform/common"
	"ivmp-service/model"
	"net/http"
	"time"
)

func init() {
	common.RegisterUpsertBeforeHook("Ai_model", ProcessUpsertAiModel)
}
func ProcessUpsertAiModel(r *http.Request, in any) (out any, err error) {
	sub, _ := common.ExtractUserSub(r)

	inObj := in.(model.Ai_model)
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
