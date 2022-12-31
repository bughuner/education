package gift

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

func UpdateGiftApi(c *gin.Context) {
	var req model.Gift
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("ShouldBindJSON gift failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	if err = checkUpdateParam(&req); err != nil {
		log.Printf("checkParam failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.NoParams, err.Error())
		return
	}
	doc, err := updateGift(c, &req)
	if err != nil {
		log.Printf("updateDoc failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.InternalServerError, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, doc)
}

func checkUpdateParam(req *model.Gift) error {
	if req.ID == "" {
		return util.BuildErrorInfo("ID为空")
	}
	return nil
}

func updateGift(c *gin.Context, req *model.Gift) (*model.Gift, error) {
	giftDb := database.Query.Gift
	gift, err := giftDb.WithContext(c).Where(giftDb.ID.Eq(req.ID)).First()
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Printf("giftDb query failed, err:%v", err)
		return nil, err
	}
	if gift == nil {
		return nil, util.BuildErrorInfo("未找到礼物")
	}
	if req.Coin != 0 {
		gift.Coin = req.Coin
	}
	if req.Name != "" {
		gift.Name = req.Name
	}
	if req.Image != "" {
		gift.Image = req.Image
	}
	if req.Introduction != "" {
		gift.Introduction = req.Introduction
	}
	gift.Count = req.Count
	err = giftDb.WithContext(c).Save(gift)
	if err != nil {
		log.Printf("giftDb save failed, err:%v", err)
		return nil, err
	}
	return gift, nil
}
