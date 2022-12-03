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

func UpdateGiftCountApi(c *gin.Context) {
	var req model.Shop
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("ShouldBindJSON shop failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	if err = checkUpdateGiftCountParam(&req); err != nil {
		log.Printf("checkParam failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.NoParams, err.Error())
		return
	}
	shop, err := UpdateShopCount(c, req.ID, req.GiftID, req.Count)
	if err != nil {
		log.Printf("UpdateShopCount failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.InternalServerError, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, shop)
}

func checkUpdateGiftCountParam(req *model.Shop) error {
	if req.ID == "" {
		return util.BuildErrorInfo("ID为空")
	}
	if req.GiftID == "" {
		return util.BuildErrorInfo("礼物ID为空")
	}
	if req.Count < 0 {
		return util.BuildErrorInfo("礼物数量不能小于0")
	}
	return nil
}

func UpdateShopCount(c *gin.Context, shopId, giftId string, count int64) (*model.Shop, error) {
	shopDb := database.Query.Shop
	shop, err := shopDb.WithContext(c).Where(shopDb.ID.Eq(shopId)).First()
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Printf("shopDb query failed, err:%v\n", err)
		return nil, util.BuildErrorInfo("shopDb query failed, err:%v", err)
	}
	if shop == nil {
		return nil, util.BuildErrorInfo("商品ID不存在")
	}
	if giftId != "" {
		shop.GiftID = giftId
	}
	if count > 0 {
		shop.Count = count
	}
	err = UpdateShopCountById(c, shopId, giftId, count)
	if err != nil {
		return nil, util.BuildErrorInfo("UpdateShopCountById failed, err:%v", err)
	}
	return shop, nil
}

func UpdateShopCountById(c *gin.Context, shopId, giftId string, count int64) error {
	shopDb := database.Query.Shop
	_, err := shopDb.WithContext(c).Where(shopDb.ID.Eq(shopId), shopDb.GiftID.Eq(giftId)).UpdateSimple(shopDb.Count.Value(count))
	if err != nil {
		log.Printf("shopDb update failed, shopId:%v, giftId:%v, count:%v", shopId, giftId, count)
		return util.BuildErrorInfo("shopDb update failed, shopId:%v, giftId:%v, count:%v", shopId, giftId, count)
	}
	return nil
}
