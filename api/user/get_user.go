package user

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

func GetUserApi(c *gin.Context) {
	var req model_view.GetUserReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("ShouldBindJSON GetUserReq failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	if err := checkGetUserParams(&req); err != nil {
		log.Printf("checkParam failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.NoParams, err.Error())
		return
	}
	userList, err := getUser(c, &req)
	if err != nil {
		log.Printf("getUser failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.OperationErr, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, userList)
}

func checkGetUserParams(req *model_view.GetUserReq) error {
	if req.PageNo == 0 {
		req.PageNo = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 100
	}
	return nil
}

func getUser(c *gin.Context, req *model_view.GetUserReq) (*model_view.GetUserResp, error) {
	userDb := database.Query.User
	sql := userDb.WithContext(c)
	if req.ID != "" {
		sql = sql.Where(userDb.ID.Eq(req.ID))
	}
	if req.Name != "" {
		sql = sql.Where(userDb.Name.Eq(req.Name))
	}
	if req.Level != 0 {
		sql = sql.Where(userDb.Level.Eq(req.Level))
	}
	if req.MapID != "" {
		sql = sql.Where(userDb.MapID.Eq(req.MapID))
	}
	if req.Account != "" {
		sql = sql.Where(userDb.Account.Eq(req.Account))
	}
	if req.Role != "" {
		sql = sql.Where(userDb.Role.Eq(req.Role))
	}
	if req.Career != "" {
		sql = sql.Where(userDb.Career.Eq(req.Career))
	}
	total, err := sql.Count()
	if err != nil {
		log.Printf("userDb count failed, err:%v\n", err)
		return nil, util.BuildErrorInfo("userDb query failed, err:%v", err)
	}
	userList, err := sql.Offset((req.PageNo - 1) * req.PageSize).Limit(req.PageSize).Find()
	if err != nil {
		log.Printf("userDb count failed, err:%v\n", err)
		return nil, util.BuildErrorInfo("userDb query failed, err:%v", err)
	}
	if userList == nil || len(userList) == 0 {
		return nil, nil
	}
	users := make([]*model.User, len(userList))
	for i, item := range userList {
		users[i] = &model.User{
			ID:               item.ID,
			Account:          item.Account,
			Name:             item.Name,
			Password:         "",
			Role:             item.Role,
			Level:            item.Level,
			Career:           item.Career,
			Image:            item.Image,
			Coin:             item.Coin,
			Credit:           item.Credit,
			MapID:            item.MapID,
			Horizon:          item.Horizon,
			Ordinate:         item.Ordinate,
			Hp:               item.Hp,
			VictoryCount:     item.VictoryCount,
			Merits:           item.Merits,
			FailCount:        item.FailCount,
			KillMonsterCount: item.KillMonsterCount,
		}
	}
	res := &model_view.GetUserResp{
		Total: total,
		Data:  users,
	}
	return res, nil
}
