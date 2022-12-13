package plat

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

func UpdatePlatApi(c *gin.Context) {
	var req model.Plat
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("ShouldBindJSON Plat failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	if err = checkUpdateParam(&req); err != nil {
		log.Printf("checkParam failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.NoParams, err.Error())
		return
	}
	monster, err := updatePlat(c, &req)
	if err != nil {
		log.Printf("updateDocQuestion failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.OperationErr, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, monster)
}

func checkUpdateParam(plat *model.Plat) error {
	if plat.ID == "" {
		return util.BuildErrorInfo("ID为空")
	}
	return nil
}

func updatePlat(c *gin.Context, req *model.Plat) (*model.Plat, error) {
	platDb := database.Query.Plat
	plat, err := platDb.WithContext(c).Where(platDb.ID.Eq(req.ID)).First()
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Printf("docQuestionDb query failed, id:%v, err:%v\n", req.ID, err)
		return nil, util.BuildErrorInfo("docQuestionDb query failed, err:%v", err)
	}
	if plat == nil {
		return nil, util.BuildErrorInfo("地图不存在")
	}
	if req.Name != "" {
		plat.Name = req.Name
	}
	if req.Image != "" {
		plat.Image = req.Image
	}
	if req.Height != 0 {
		plat.Height = req.Height
	}
	if req.Width != 0 {
		plat.Width = req.Width
	}
	if req.PassArea != "" {
		plat.PassArea = req.PassArea
	}
	err = platDb.WithContext(c).Save(plat)
	if err != nil {
		log.Printf("platDb save failed, oldPlat:%v, err:%v", plat, err)
		return nil, util.BuildErrorInfo("platDb save failed, err:%v", err)
	}
	return plat, nil
}
