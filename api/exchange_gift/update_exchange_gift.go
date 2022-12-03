package exchange_gift

import (
	"education/common"
	errno "education/common/erron"
	"education/database"
	"education/model"
	"education/util"
	"github.com/gin-gonic/gin"
	"log"
)

func UpdateExchangeGiftApi(c *gin.Context) {
	var req model.ExchangeGift
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("ShouldBindJSON exchange_gift failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	if err := checkUpdateParams(&req); err != nil {
		log.Printf("checkParam failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.NoParams, err.Error())
		return
	}
	exchange, err := updateExchangeGift(c, req.ID, req.UserID, req.GiftID, req.Count, req.IsExchange)
	if err != nil {
		log.Printf("updateExchangeGift failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.InternalServerError, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, exchange)
}

func checkUpdateParams(req *model.ExchangeGift) error {
	if req.ID == "" {
		return util.BuildErrorInfo("ID 不能为空")
	}
	if req.Count < 0 {
		return util.BuildErrorInfo("个数不能小于0")
	}
	return nil
}

func updateExchangeGift(c *gin.Context, id, userId, giftId string, count, isExchange int64) (*model.ExchangeGift, error) {
	exchangeGiftDb := database.Query.ExchangeGift
	exchange, err := exchangeGiftDb.WithContext(c).Where(exchangeGiftDb.ID.Eq(id)).First()
	if err != nil {
		log.Printf("exchangeGiftDb query failed, err:%v\n", err)
		return nil, util.BuildErrorInfo("exchangeGiftDb query failed, err:%v", err)
	}
	if exchange == nil {
		return nil, util.BuildErrorInfo("记录不存在")
	}
	if userId != "" {
		exchange.UserID = userId
	}
	if giftId != "" {
		exchange.GiftID = giftId
	}
	if count >= 0 {
		exchange.Count = count
	}
	if isExchange >= 0 {
		exchange.IsExchange = isExchange
	}
	err = exchangeGiftDb.WithContext(c).Save(exchange)
	if err != nil {
		log.Printf("exchangeGiftDb save failed, err:%v\n", err)
		return nil, util.BuildErrorInfo("exchangeGiftDb save failed, err:%v", err)
	}
	return exchange, nil
}
