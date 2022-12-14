package form

import (
	"education/common"
	errno "education/common/erron"
	"education/database"
	"education/model/model_view"
	"education/util"
	"github.com/gin-gonic/gin"
	"log"
)

func GetFormApi(c *gin.Context) {
	var req model_view.GetFormReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("ShouldBindJSON form failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	if err = checkGetParam(&req); err != nil {
		log.Printf("checkParam failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.NoParams, err.Error())
		return
	}
	res, err := getForm(c, &req)
	if err != nil {
		log.Printf("getForm failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.OperationErr, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, res)
}

func checkGetParam(req *model_view.GetFormReq) error {
	if req.PageNo == 0 {
		req.PageNo = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 100
	}
	return nil
}

func getForm(c *gin.Context, req *model_view.GetFormReq) (*model_view.GetFormResp, error) {
	formDb := database.Query.Form
	sql := formDb.WithContext(c)
	if req.ID != "" {
		sql = sql.Where(formDb.ID.Eq(req.ID))
	}
	if req.UserID != "" {
		sql = sql.Where(formDb.UserID.Eq(req.UserID))
	}
	if req.Source != 0 {
		sql = sql.Where(formDb.Source.Eq(req.Source))
	}
	if req.MapID != "" {
		sql = sql.Where(formDb.MapID.Eq(req.MapID))
	}
	if req.Type != 0 {
		sql = sql.Where(formDb.Type.Eq(req.Type))
	}
	total, err := sql.Count()
	if err != nil {
		log.Printf("formDb count failed, err:%v\n", err)
		return nil, util.BuildErrorInfo("formDb count failed, err:%v", err)
	}
	formList, err := sql.Offset((req.PageNo - 1) * req.PageSize).Limit(req.PageSize).Find()
	if err != nil {
		log.Printf("formDb query failed, err:%v\n", err)
		return nil, util.BuildErrorInfo("formDb query failed, err:%v\n", err)
	}
	res := &model_view.GetFormResp{
		Total: total,
		Data:  formList,
	}
	return res, nil
}
