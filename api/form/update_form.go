package form

import (
	"education/common"
	errno "education/common/erron"
	"education/database"
	"education/model"
	"education/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

func UpdateFormApi(c *gin.Context) {
	var req model.Form
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("ShouldBindJSON form failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	if err = checkUpdateParam(&req); err != nil {
		log.Printf("checkParam failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.NoParams, err.Error())
		return
	}
	form, err := updateForm(c, &req)
	if err != nil {
		log.Printf("updateForm failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.OperationErr, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, form)
}

func checkUpdateParam(req *model.Form) error {
	if req.ID == "" {
		return util.BuildErrorInfo("ID为空")
	}
	return nil
}

func updateForm(c *gin.Context, req *model.Form) (*model.Form, error) {
	formDb := database.Query.Form
	oldForm, err := formDb.WithContext(c).Where(formDb.ID.Eq(req.ID)).First()
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Printf("formDb query failed, req:%v, err:%v\n", req, err)
		return nil, util.BuildErrorInfo("formDb query failed, err:%v", err)
	}
	if oldForm == nil {
		return nil, util.BuildErrorInfo("通知不存在")
	}
	if req.Source != 0 {
		oldForm.Source = req.Source
	}
	if req.Content != "" {
		oldForm.Content = req.Content
	}
	if req.Source != 0 {
		oldForm.Source = req.Source
	}
	if req.MapID != "" {
		oldForm.MapID = req.MapID
	}
	if req.UserID != "" {
		oldForm.UserID = req.UserID
	}
	if req.Type != 0 {
		oldForm.Type = req.Type
	}
	err = formDb.WithContext(c).Save(oldForm)
	if err != nil {
		log.Printf("docQuestionDb save failed, oldForm:%v, err:%v", oldForm, err)
		return nil, util.BuildErrorInfo("docQuestionDb save failed, err:%v", err)
	}
	return oldForm, nil
}
