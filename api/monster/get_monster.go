package monster

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

func GetMonsterApi(c *gin.Context) {
	var req model_view.GetMonsterReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("ShouldBindJSON GetMonsterReq failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	if err := checkGetParam(&req); err != nil {
		log.Printf("checkParam failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.NoParams, err.Error())
		return
	}
	monster, err := getMonsterInfo(c, &req)
	if err != nil {
		log.Printf("getMonsterInfo failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.OperationErr, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, monster)
}

func checkGetParam(req *model_view.GetMonsterReq) error {
	if req.PageNo == 0 {
		req.PageNo = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 1
	}
	return nil
}

func getMonsterInfo(c *gin.Context, req *model_view.GetMonsterReq) (*model_view.GetMonsterResp, error) {
	monsterDb := database.Query.Monster
	sql := monsterDb.WithContext(c)
	if req.ID != "" {
		sql = sql.Where(monsterDb.ID.Eq(req.ID))
	}
	if req.Name != "" {
		sql = sql.Where(monsterDb.Name.Eq(req.Name))
	}
	if req.Type != 0 {
		sql = sql.Where(monsterDb.Type.Eq(req.Type))
	}
	if req.Hp != 0 {
		sql = sql.Where(monsterDb.Hp.Eq(req.Hp))
	}
	if req.Level != 0 {
		sql = sql.Where(monsterDb.Level.Eq(req.Level))
	}
	if req.Coin != 0 {
		sql = sql.Where(monsterDb.Coin.Eq(req.Coin))
	}
	if req.ExperienceValue != 0 {
		sql = sql.Where(monsterDb.ExperienceValue.Eq(req.ExperienceValue))
	}
	total, err := sql.Count()
	if err != nil {
		log.Printf("monsterDb count failed, err:%v\n", err)
		return nil, util.BuildErrorInfo("monsterDb count failed, err:%v", err)
	}
	monsterList, err := sql.Offset(req.PageNo - 1).Limit(req.PageSize).Find()
	if err != nil {
		log.Printf("monsterDb query failed, err:%v\n", err)
		return nil, util.BuildErrorInfo("monsterDb query failed, err:%v\n", err)
	}
	monsters := make([]*model.Monster, len(monsterList))
	for i, item := range monsterList {
		monsters[i] = &model.Monster{
			ID:              item.ID,
			Type:            item.Type,
			Name:            item.Name,
			Level:           item.Level,
			ExperienceValue: item.ExperienceValue,
			Coin:            item.Coin,
			Hp:              item.Hp,
			Image:           item.Image,
		}
	}
	res := &model_view.GetMonsterResp{
		Total: total,
		Data:  monsters,
	}
	return res, nil
}
