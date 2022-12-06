package form

import (
	"education/common"
	errno "education/common/erron"
	"education/database"
	"education/model"
	"education/model/model_view"
	"education/util"
	"github.com/gin-gonic/gin"
	"log"
)

func DeleteFormApi(c *gin.Context) {
	var req model_view.DeleteFromReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("ShouldBindJSON form failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	err = deleteForm(c, req.FormIds)
	if err != nil {
		log.Printf("deleteDocQuestion failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.OperationErr, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, nil)
}

func deleteForm(c *gin.Context, ids []string) error {
	formDb := database.Query.Form
	_, err := formDb.WithContext(c).Where(formDb.ID.In(ids...)).Delete(&model.Form{})
	if err != nil {
		log.Printf("formDb delete failed, err:%v", err)
		return util.BuildErrorInfo("formDb delete failed, err:%v", err)
	}
	return nil
}
