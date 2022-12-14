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
	total             int64
}

// 获取奖品详情
func GetGiftCountApi(c *gin.Context) {
	var req model_view.GetGiftCountReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("ShouldBindJSON GetGiftCountReq failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	if err := checkGetGiftCountParams(&req); err != nil {
		log.Printf("checkParam failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.NoParams, err.Error())
		return
	}
	shopGiftCountList, err := getGiftCount(c, &req)
	if err != nil {
		log.Printf("getGiftCount failed, req:%v, err:%v", req, err)
		common.SendResponse(c, errno.OperationErr, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, shopGiftCountList)
}

func checkGetGiftCountParams(req *model_view.GetGiftCountReq) error {
	if req.PageNo == 0 {
		req.PageNo = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 100
	}
	return nil
}

func getGiftCount(c *gin.Context, req *model_view.GetGiftCountReq) (*model_view.GetGiftCountResp, error) {
	shopDb := database.Query.Shop
	sql := shopDb.WithContext(c)
	if req.ID != "" {
		sql = sql.Where(shopDb.ID.Eq(req.ID))
	}
	if req.GiftID != "" {
		sql = sql.Where(shopDb.GiftID.Eq(req.GiftID))
	}
	total, err := sql.Count()
	if err != nil {
		log.Printf("shopDb count failed, req:%v, err:%v\n", req, err)
		return nil, util.BuildErrorInfo("shopDb count failed, req:%v, err:%v", req, err)
	}
	shopList, err := sql.Offset((req.PageNo - 1) * req.PageSize).Limit(req.PageSize).Find()
	if err != nil {
		log.Printf("shopDb query failed, req:%v, err:%v\n", req, err)
		return nil, util.BuildErrorInfo("shopDb query failed, req:%v, err:%v\n", req, err)
	}
	giftIds := make([]string, len(shopList))
	for i, item := range shopList {
		giftIds[i] = item.GiftID
	}
	giftList, err := gift.GetGiftById(c, giftIds)
	if err != nil {
		log.Printf("GetGiftById failed, giftIds:%v, err:%v\n", giftIds, err)
		return nil, util.BuildErrorInfo("GetGiftById failed, giftIds:%v, err:%v", giftIds, err)
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
	res := &model_view.GetGiftCountResp{
		Total: total,
		Data:  shopGiftCount,
	}
	return res, nil
}
