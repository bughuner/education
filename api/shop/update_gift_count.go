package shop

import (
	"education/database"
	"education/util"
	"github.com/gin-gonic/gin"
	"log"
)

//func UpdateGiftCountApi(c *gin.Context) {
//	var req model.Shop
//	err := c.ShouldBindJSON(&req)
//	if err != nil {
//		log.Printf("ShouldBindJSON shop failed, err:%v\n", err)
//		common.SendResponse(c, errno.ErrParams, err.Error())
//		return
//	}
//	if err = checkUpdateGiftCountParam(&req); err != nil {
//		log.Printf("checkParam failed, req:%v, err:%v\n", req, err)
//		common.SendResponse(c, errno.NoParams, err.Error())
//		return
//	}
//	err = UpdateShopCountById(c, req.ID, req.GiftID, req.Count)
//	if err != nil {
//
//	}
//}

//func checkUpdateGiftCountParam(req *model.Shop) error {
//	if req.ID == "" {
//		return util.BuildErrorInfo("ID为空")
//	}
//	if req.GiftID == "" {
//		return util.BuildErrorInfo("礼物ID为空")
//	}
//	if req.Count < 0 {
//		return util.BuildErrorInfo("礼物数量不能小于0")
//	}
//	return nil
//}


func UpdateShopCountById(c *gin.Context, shopId, giftId string, count int64) error {
	shopDb := database.Query.Shop
	_, err := shopDb.WithContext(c).Where(shopDb.ID.Eq(shopId), shopDb.GiftID.Eq(giftId)).UpdateSimple(shopDb.Count.Value(count))
	if err != nil {
		log.Printf("shopDb update failed, shopId:%v, giftId:%v, count:%v", shopId, giftId, count)
		return util.BuildErrorInfo("shopDb update failed, shopId:%v, giftId:%v, count:%v", shopId, giftId, count)
	}
	return nil
}