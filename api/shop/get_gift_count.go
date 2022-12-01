package shop

import (
	"education/api/gift"
	"education/common"
	errno "education/common/erron"
	"education/database"
	"education/model"
	"education/model/model_view"
	"education/util"
	"github.com/gin-gonic/gin"
	"log"
)

type ShopGiftCountInfo struct {
	ShopGiftCountList []*model_view.ShopGiftCount
	total int64
}

func GetGiftCountApi(c *gin.Context) {
	pageNo := c.Query("page_no")
	pageSize := c.Query("page_size")
	pageNoInt, err := util.ConvertStringToInt(pageNo)
	if err != nil {
		log.Printf("数据转换错误")
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	pageSizeInt, err :=  util.ConvertStringToInt(pageSize)
	if err != nil {
		log.Printf("数据转换错误")
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	if pageNoInt == 0 || pageSizeInt == 0 {
		pageNoInt = 1
		pageSizeInt = 100
	}
	shopGiftCountList,total, err := getGiftCount(c, pageNoInt, pageSizeInt)
	if err != nil {
		log.Printf("getGiftCount failed, pageNo:%v, pageSize:%v, err:%v", pageNoInt, pageSizeInt, err)
		common.SendResponse(c, errno.OperationErr, err.Error())
		return
	}
	shopGiftCountInfo := &ShopGiftCountInfo{
		ShopGiftCountList: shopGiftCountList,
		total:             total,
	}
	common.SendResponse(c, errno.OK, shopGiftCountInfo)
}

func getGiftCount(c *gin.Context, pageNo, pageSize int)([]*model_view.ShopGiftCount, int64, error) {
	shopDb := database.Query.Shop
	sql := shopDb.WithContext(c)
	total, err := sql.Count()
	if err != nil {
		log.Printf("shopDb count failed, pageNo:%v, pageSize:%v, err:%v", pageNo, pageSize, err)
		return nil, 0, util.BuildErrorInfo("shopDb count failed, pageNo:%v, pageSize:%v, err:%v", pageNo, pageSize, err)
	}
	shopList, err := sql.Offset(pageNo - 1).Limit(pageSize).Find()
	if err != nil {
		log.Printf("shopDb query failed, pageNo:%v, pageSize:%v, err:%v", pageNo, pageSize, err)
		return nil, 0, util.BuildErrorInfo("shopDb query failed, pageNo:%v, pageSize:%v, err:%v", pageNo, pageSize, err)
	}
	giftIds := make([]string, len(shopList))
	for i, item :=  range shopList {
		giftIds[i] = item.GiftID
	}
	giftList, err := gift.GetGiftById(c, giftIds)
	if err != nil {
		log.Printf("GetGiftById failed, pageNo:%v, pageSize:%v, err:%v", pageNo, pageSize, err)
		return nil, 0, util.BuildErrorInfo("shopDb GetGiftById failed, pageNo:%v, pageSize:%v, err:%v", pageNo, pageSize, err)
	}
	giftListMap := make(map[string]*model.Gift)
	for _, item := range giftList {
		giftListMap[item.ID] = item
	}
	shopGiftCount := make([]*model_view.ShopGiftCount, len(shopList))
	for i, item := range shopList {
		gift, ok := giftListMap[item.GiftID]
		if !ok {
			log.Printf("gift not exist, giftId:%v", item.GiftID)
			continue
		}
		shopGiftCount[i] = &model_view.ShopGiftCount{
			ID:           item.ID,
			GiftID:       item.GiftID,
			Count:        item.Count,
			Name:         gift.Name,
			Introduction: gift.Introduction,
			Image:        gift.Image,
			Coin:         gift.Coin,
		}
	}
	return shopGiftCount, total, nil
}
