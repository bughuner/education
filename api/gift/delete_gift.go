package gift

import (
	"education/common"
	errno "education/common/erron"
	"education/database"
	"education/model"
	"education/model/model_view"
	"education/util"
	"github.com/gin-gonic/gin"
	"log"
)

func DeleteGiftApi(c *gin.Context) {
	var req model_view.DeleteGiftReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("ShouldBindJSON DeleteGiftReq failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	err = deleteGift(c, req.GiftIds)
	if err != nil {
		log.Printf("deleteGift failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.OperationErr, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, nil)
}

func deleteGift(c *gin.Context, ids []string) error {
	giftDb := database.Query.Gift
	_, err := giftDb.WithContext(c).Where(giftDb.ID.In(ids...)).Delete(&model.Gift{})
	if err != nil {
		log.Printf("giftDb delete failed, err:%v", err)
		return util.BuildErrorInfo("giftDb delete failed, err:%v", err)
	}
	return nil
}
