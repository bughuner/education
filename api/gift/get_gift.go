package gift

import (
	"education/database"
	"education/model"
	"education/util"
	"github.com/gin-gonic/gin"
	"log"
)

func GetGiftById(c *gin.Context, ids []string) ([]*model.Gift, error) {
	giftDb := database.Query.Gift
	giftList, err := giftDb.WithContext(c).Where(giftDb.ID.In(ids...)).Find()
	if err != nil {
		log.Printf("giftDb query failed, ids:%v, err:%v\n", ids, err)
		return nil, util.BuildErrorInfo("giftDb query failed, ids:%v, err:%v", ids, err)
	}
	return giftList, nil
}