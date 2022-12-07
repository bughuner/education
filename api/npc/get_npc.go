package npc

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

func GetNpcApi(c *gin.Context) {
	var req model_view.GetNpcReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("ShouldBindJSON GetNpcReq failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	if err := checkGetNPCParam(&req); err != nil {
		log.Printf("checkParam failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.NoParams, err.Error())
		return
	}
	taskList, err := getNpc(c, &req)
	if err != nil {
		log.Printf("getNpcTask failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.OperationErr, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, taskList)
}

func checkGetNPCParam(req *model_view.GetNpcReq) error {
	if req.PageNo == 0 {
		req.PageNo = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 100
	}
	return nil
}

func getNpc(c *gin.Context, req *model_view.GetNpcReq) (*model_view.GetNpcResp, error) {
	npcDb := database.Query.Npc
	sql := npcDb.WithContext(c)
	if req.ID != "" {
		sql = sql.Where(npcDb.ID.Eq(req.ID))
	}
	if req.Name != "" {
		sql = sql.Where(npcDb.Name.Eq(req.Name))
	}
	if req.MapID != "" {
		sql = sql.Where(npcDb.MapID.Eq(req.MapID))
	}
	if req.Type != 0 {
		sql = sql.Where(npcDb.Type.Eq(req.Type))
	}
	total, err := sql.Count()
	if err != nil {
		log.Printf("npcDb count failed, err:%v\n", err)
		return nil, util.BuildErrorInfo("npcDb count failed, err:%v", err)
	}
	npcList, err := sql.Offset(req.PageNo - 1).Limit(req.PageSize).Find()
	if err != nil {
		log.Printf("npcDb query failed, err:%v\n", err)
		return nil, util.BuildErrorInfo("npcDb query failed, err:%v\n", err)
	}
	npcs := make([]*model.Npc, len(npcList))
	for i, item := range npcList {
		npcs[i] = &model.Npc{
			ID:           item.ID,
			Name:         item.Name,
			Image:        item.Image,
			Introduction: item.Introduction,
			Horizon:      item.Horizon,
			Ordinate:     item.Ordinate,
			MapID:        item.MapID,
			Type:         item.Type,
		}
	}
	res := &model_view.GetNpcResp{
		Total: total,
		Data:  npcs,
	}
	return res, nil
}
