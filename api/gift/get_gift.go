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

func GetGiftApi(c *gin.Context) {
	var req model_view.GetGiftReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("ShouldBindJSON GetGiftReq failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	if err = checkGetParam(&req); err != nil {
		log.Printf("checkParam failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.NoParams, err.Error())
		return
	}
	res, err := getGift(c, &req)
	if err != nil {
		log.Printf("getGift failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.OperationErr, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, res)
}

func checkGetParam(req *model_view.GetGiftReq) error {
	if req.PageNo == 0 {
		req.PageNo = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 100
	}
	return nil
}

func getGift(c *gin.Context, req *model_view.GetGiftReq) (*model_view.GetGiftResp, error) {
	giftDb := database.Query.Gift
	sql := giftDb.WithContext(c)
	if req.ID != "" {
		sql = sql.Where(giftDb.ID.Eq(req.ID))
	}
	if req.Name != "" {
		sql = sql.Where(giftDb.Name.Eq(req.Name))
	}
	if req.Coin != 0 {
		sql = sql.Where(giftDb.Coin.Eq(req.Coin))
	}
	total, err := sql.Count()
	if err != nil {
		log.Printf("giftDb count failed, err:%v\n", err)
		return nil, util.BuildErrorInfo("giftDb count failed, err:%v", err)
	}
	giftList, err := sql.Offset((req.PageNo - 1) * req.PageSize).Limit(req.PageSize).Find()
	if err != nil {
		log.Printf("giftDb query failed, err:%v\n", err)
		return nil, util.BuildErrorInfo("giftDb query failed, err:%v\n", err)
	}
	res := &model_view.GetGiftResp{
		Total: total,
		Data:  giftList,
	}
	return res, nil
}

func GetGiftById(c *gin.Context, ids []string) ([]*model.Gift, error) {
	giftDb := database.Query.Gift
	giftList, err := giftDb.WithContext(c).Where(giftDb.ID.In(ids...)).Find()
	if err != nil {
		log.Printf("giftDb query failed, ids:%v, err:%v\n", ids, err)
		return nil, util.BuildErrorInfo("giftDb query failed, ids:%v, err:%v", ids, err)
	}
	return giftList, nil
}
