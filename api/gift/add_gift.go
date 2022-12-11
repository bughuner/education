package gift

import (
	"education/common"
	errno "education/common/erron"
	"education/database"
	"education/model"
	"education/util"
	"github.com/gin-gonic/gin"
	"log"
)

func AddGiftApi(c *gin.Context) {
	var req model.Gift
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("ShouldBindJSON gift failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	if err = checkAddParam(&req); err != nil {
		log.Printf("checkParam failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.NoParams, err.Error())
		return
	}
	gift, err := addGift(c, &req)
	if err != nil {
		log.Printf("addGift failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.InternalServerError, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, gift)
}

func checkAddParam(gift *model.Gift) error {
	if gift.Name == "" {
		return util.BuildErrorInfo("名称为空")
	}
	if gift.Coin == 0 {
		return util.BuildErrorInfo("价值为0")
	}
	return nil
}

func addGift(c *gin.Context, gift *model.Gift) (*model.Gift, error) {
	giftDb := database.Query.Gift
	id := util.GetUUID()
	gift.ID = id
	err := giftDb.WithContext(c).Create(gift)
	if err != nil {
		log.Printf("giftDb create failed, err:%v", err)
		return nil, err
	}
	return gift, nil
}
