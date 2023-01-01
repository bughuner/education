package exchange_gift

import (
	"education/common"
	errno "education/common/erron"
	"education/database"
	"education/model"
	"education/query"
	"education/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"time"
)

func ExchangeGiftApi(c *gin.Context) {
	var req model.ExchangeGift
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("ShouldBindJSON exchange_gift failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	if err := checkParam(&req); err != nil {
		log.Printf("checkParam failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.NoParams, err.Error())
		return
	}
	err = exchangeGift(c, req.UserID, req.GiftID, req.Count)
	if err != nil {
		log.Printf("exchangeGift failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.OperationErr, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, nil)
}

func checkParam(exchangeGift *model.ExchangeGift) error {
	if exchangeGift.GiftID == "" {
		return util.BuildErrorInfo("礼物ID为空")
	}
	if exchangeGift.UserID == "" {
		return util.BuildErrorInfo("用户ID为空")
	}
	if exchangeGift.Count <= 0 {
		return util.BuildErrorInfo("兑换数量需要大于0")
	}
	return nil
}

func exchangeGift(c *gin.Context, userId, giftId string, count int64) error {
	lockKey := fmt.Sprintf("gift_%v", giftId)
	lockValue := time.Now().UnixNano()
	lockOk := database.SetLock(c, lockKey, lockValue, 3*time.Second)
	defer database.DelLock(c, lockKey, lockValue)
	if !lockOk {
		log.Printf("获取锁冲突\n")
		return util.BuildErrorInfo("获取锁冲突")
	}
	giftDb := database.Query.Gift
	gift, err := giftDb.WithContext(c).Where(giftDb.ID.Eq(giftId)).First()
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Printf("shopDb query failed, err:%v\n", err)
		return util.BuildErrorInfo("shopDb query failed, err:%v", err)
	}
	if gift == nil {
		log.Printf("gift is nil\n")
		return util.BuildErrorInfo("商品不存在")
	}
	if gift.Count < count {
		log.Printf("礼物存量不足\n")
		return util.BuildErrorInfo("礼物存量不足")
	}
	userDb := database.Query.User
	user, err := userDb.WithContext(c).Where(userDb.ID.Eq(userId)).First()
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Printf("userDb query failed, err:%v\n", err)
		return util.BuildErrorInfo("userDb query failed, err:%v", err)
	}
	if user == nil {
		log.Printf("user is nil\n")
		return util.BuildErrorInfo("用户不存在")
	}
	if user.Coin < count*gift.Coin {
		log.Printf("用户钱不够\n")
		return util.BuildErrorInfo("用户钱不够")
	}
	exchangeGift := &model.ExchangeGift{
		ID:     util.GetUUID(),
		UserID: userId,
		GiftID: giftId,
		Count:  count,
		Time:   time.Now().UnixNano() / 1e6, // 毫秒
	}
	database.Query.Transaction(func(tx *query.Query) error {
		_, err = tx.Gift.WithContext(c).Where(tx.Gift.ID.Eq(giftId)).UpdateSimple(tx.Gift.Count.Value(gift.Count - count))
		if err != nil {
			log.Printf("shopDb update failed, err:%v\n", err)
			return util.BuildErrorInfo("shopDb update failed, err:%v", err)
		}
		_, err = tx.User.WithContext(c).Where(tx.User.ID.Eq(userId)).UpdateSimple(tx.User.Coin.Value(user.Coin - count*gift.Coin))
		if err != nil {
			log.Printf("User update failed, err:%v\n", err)
			return util.BuildErrorInfo("User update failed, err:%v", err)
		}
		err = tx.ExchangeGift.WithContext(c).Create(exchangeGift)
		if err != nil {
			log.Printf("shopDb update failed, err:%v\n", err)
			return util.BuildErrorInfo("shopDb update failed, err:%v", err)
		}
		return nil
	})
	return nil
}
