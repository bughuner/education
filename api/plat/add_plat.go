package plat

import (
	"education/common"
	errno "education/common/erron"
	"education/database"
	"education/model"
	"education/util"
	"github.com/gin-gonic/gin"
	"log"
)

func AddPlatApi(c *gin.Context) {
	var req model.Plat
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("ShouldBindJSON plat failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	if err = checkAddParam(&req); err != nil {
		log.Printf("checkParam failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.NoParams, err.Error())
		return
	}
	plat, err := addPlat(c, &req)
	if err != nil {
		log.Printf("addPlat failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.OperationErr, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, plat)
}

func checkAddParam(req *model.Plat) error {
	if req.Name == "" {
		return util.BuildErrorInfo("名称为空")
	}
	return nil
}

func addPlat(c *gin.Context, plat *model.Plat) (*model.Plat, error) {
	id := util.GetUUID()
	plat.ID = id
	platDb := database.Query.Plat
	err := platDb.WithContext(c).Create(plat)
	if err != nil {
		log.Printf("platDb create failed, err:%v\n", err)
		return nil, util.BuildErrorInfo("platDb create failed, err:%v", err)
	}
	return plat, nil
}
