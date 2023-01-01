package exchange_gift

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

func GetExchangeGiftApi(c *gin.Context) {
	var req model_view.ExchangeGiftReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("ShouldBindJSON exchange_gift failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	if err := checkGetParams(&req); err != nil {
		log.Printf("checkParam failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.NoParams, err.Error())
		return
	}
	res, err := getExchangeGiftList(c, &req)
	if err != nil {
		log.Printf("getExchangeGiftList failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.InternalServerError, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, res)
}

func checkGetParams(req *model_view.ExchangeGiftReq) error {
	if req.PageNo == 0 {
		req.PageNo = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 100
	}
	return nil
}

func getExchangeGiftList(c *gin.Context, req *model_view.ExchangeGiftReq) (*model_view.ExchangeGiftResp, error) {
	exchangeGiftDb := database.Query.ExchangeGift
	sql := exchangeGiftDb.WithContext(c)
	if req.UserID != "" {
		sql = sql.Where(exchangeGiftDb.UserID.Eq(req.UserID))
	}
	if req.GiftID != "" {
		sql = sql.Where(exchangeGiftDb.GiftID.Eq(req.GiftID))
	}
	if req.Count != 0 {
		sql = sql.Where(exchangeGiftDb.Count.Eq(req.Count))
	}
	total, err := sql.Count()
	if err != nil {
		log.Printf("exchangeGiftDb get total failed, err:%v\n", err)
		return nil, util.BuildErrorInfo("exchangeGiftDb get total failed, err:%v", err)
	}
	exchangeGiftList, err := sql.Offset((req.PageNo - 1) * req.PageSize).Limit(req.PageSize).Find()
	if err != nil {
		log.Printf("exchangeGiftDb query failed, err:%v\n", err)
		return nil, util.BuildErrorInfo("exchangeGiftDb query failed, err:%v", err)
	}
	giftIds := make([]string, len(exchangeGiftList))
	for i, item := range exchangeGiftList {
		giftIds[i] = item.GiftID
	}
	giftList, err := gift.GetGiftById(c, giftIds)
	if err != nil {
		log.Printf("gift GetGiftById failed, err:%v\n", err)
		return nil, util.BuildErrorInfo("gift GetGiftById failed, err:%v", err)
	}
	giftMap := make(map[string]*model.Gift)
	for _, item := range giftList {
		giftMap[item.ID] = item
	}
	exchangeGiftAndGift := make([]*model_view.ExchangeGiftAndGift, len(exchangeGiftList))
	for i, item := range exchangeGiftList {
		exchangeGiftAndGift[i] = &model_view.ExchangeGiftAndGift{
			ExchangeGift: item,
			Gift:         giftMap[item.GiftID],
		}
	}
	res := &model_view.ExchangeGiftResp{
		Data:  exchangeGiftAndGift,
		Total: total,
	}
	return res, nil
}
