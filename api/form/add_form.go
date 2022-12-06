package form

import (
	"education/common"
	errno "education/common/erron"
	"education/database"
	"education/model"
	"education/util"
	"github.com/gin-gonic/gin"
	"log"
)

func AddFormApi(c *gin.Context) {
	var req model.Form
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("ShouldBindJSON form failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	if err = checkAddParam(&req); err != nil {
		log.Printf("checkParam failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.NoParams, err.Error())
		return
	}
	form, err := addForm(c, &req)
	if err != nil {
		log.Printf("addForm failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.OperationErr, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, form)
}

func checkAddParam(req *model.Form) error {
	if req.UserID == "" {
		return util.BuildErrorInfo("用户ID为空")
	}
	if req.Type == 0 {
		return util.BuildErrorInfo("未知通知类型")
	}
	if req.Content == "" {
		return util.BuildErrorInfo("通知内容为空")
	}
	if req.MapID == "" {
		return util.BuildErrorInfo("地图ID为空")
	}
	if req.Source == 0 {
		return util.BuildErrorInfo("通知来源未知")
	}
	return nil
}

func addForm(c *gin.Context, form *model.Form) (*model.Form, error) {
	id := util.GetUUID()
	form.ID = id
	formDb := database.Query.Form
	err := formDb.WithContext(c).Create(form)
	if err != nil {
		log.Printf("formDb create failed, err:%v\n", err)
		return nil, util.BuildErrorInfo("formDb create failed, err:%v", err)
	}
	return form, nil
}
