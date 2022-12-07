package plat

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

func GetPlatApi(c *gin.Context) {
	var req model_view.GetPlatReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("ShouldBindJSON map failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	if err = checkGetParam(&req); err != nil {
		log.Printf("checkParam failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.NoParams, err.Error())
		return
	}
	mapList, err := getMap(c, &req)
	if err != nil {
		log.Printf("getMap failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.OperationErr, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, mapList)
}

func checkGetParam(req *model_view.GetPlatReq) error {
	if req.PageNo == 0 {
		req.PageNo = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 100
	}
	return nil
}

func getMap(c *gin.Context, req *model_view.GetPlatReq) (*model_view.GetPlatResp, error) {
	platDb := database.Query.Plat
	sql := platDb.WithContext(c)
	if req.ID != "" {
		sql = sql.Where(platDb.ID.Eq(req.ID))
	}
	if req.Name != "" {
		sql = sql.Where(platDb.Name.Eq(req.Name))
	}
	total, err := sql.Count()
	if err != nil {
		log.Printf("platDb count failed, err:%v\n", err)
		return nil, util.BuildErrorInfo("platDb count failed, err:%v", err)
	}
	platList, err := sql.Offset(req.PageNo - 1).Limit(req.PageSize).Find()
	if err != nil {
		log.Printf("platDb query failed, err:%v\n", err)
		return nil, util.BuildErrorInfo("platDb query failed, err:%v\n", err)
	}
	plats := make([]*model.Plat, len(platList))
	for i, item := range platList {
		plats[i] = &model.Plat{
			ID:     item.ID,
			Name:   item.Name,
			Image:  item.Image,
			Width:  item.Width,
			Height: item.Height,
		}
	}
	res := &model_view.GetPlatResp{
		Total: total,
		Data:  plats,
	}
	return res, nil
}
