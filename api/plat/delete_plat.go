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

func DeletePlatApi(c *gin.Context) {
	var req model_view.DeletePlatReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("ShouldBindJSON DeletePlatReq failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	err = deletePlat(c, req.PlatIds)
	if err != nil {
		log.Printf("deleteDocQuestion failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.OperationErr, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, nil)
}

func deletePlat(c *gin.Context, ids []string) error {
	platDb := database.Query.Plat
	_, err := platDb.WithContext(c).Where(platDb.ID.In(ids...)).Delete(&model.Plat{})
	if err != nil {
		log.Printf("platDb delete failed, err:%v", err)
		return util.BuildErrorInfo("platDb delete failed, err:%v", err)
	}
	return nil
}
