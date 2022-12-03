package shop

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

// 发布奖品
func AddGiftCountApi(c *gin.Context) {
	var req model.Shop
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("ShouldBindJSON shop failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	if err = checkAddGiftCountParam(&req); err != nil {
		log.Printf("checkParam failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.NoParams, err.Error())
		return
	}
	shop, err := addGift(c, req.GiftID, req.Count)
	if err != nil {
		log.Printf("addGift failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.InternalServerError, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, shop)
}

func checkAddGiftCountParam(req *model.Shop) error {
	if req.GiftID == "" {
		return util.BuildErrorInfo("礼物ID为空")
	}
	if req.Count <= 0 {
		return util.BuildErrorInfo("礼物数量不能小于等于0")
	}
	return nil
}

func addGift(c *gin.Context, giftId string, count int64) (*model.Shop, error) {
	shopDb := database.Query.Shop
	oldShop, err := shopDb.WithContext(c).Where(shopDb.GiftID.Eq(giftId)).First()
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Printf("shopDb query failed, err:%v\n", err)
		return nil, util.BuildErrorInfo("shopDb query failed, err:%v", err)
	}
	if oldShop != nil {
		return UpdateShopCount(c, oldShop.ID, giftId, oldShop.Count+count)
	}
	id := util.GetUUID()
	gift := &model.Shop{
		ID:     id,
		GiftID: giftId,
		Count:  count,
	}
	err = shopDb.WithContext(c).Create(gift)
	if err != nil {
		log.Printf("shopDb create failed, err:%v\n", err)
		return nil, util.BuildErrorInfo("shopDb create failed, err:%v", err)
	}
	return gift, nil
}
