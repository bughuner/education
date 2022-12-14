package npc_task

import (
	"education/common"
	errno "education/common/erron"
	"education/database"
	"education/model/model_view"
	"education/util"
	"github.com/gin-gonic/gin"
	"log"
)

func GetNPCTaskApi(c *gin.Context) {
	var req model_view.GetNpcTaskReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("ShouldBindJSON GetNpcTaskReq failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	if err := checkGetNPCTaskParam(&req); err != nil {
		log.Printf("checkParam failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.NoParams, err.Error())
		return
	}
	taskList, err := getNpcTask(c, &req)
	if err != nil {
		log.Printf("getNpcTask failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.OperationErr, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, taskList)
}

func checkGetNPCTaskParam(req *model_view.GetNpcTaskReq) error {
	if req.PageNo == 0 {
		req.PageNo = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 100
	}
	return nil
}

func getNpcTask(c *gin.Context, req *model_view.GetNpcTaskReq) (*model_view.GetNpcTaskResp, error) {
	npcTaskDb := database.Query.NpcTask
	sql := npcTaskDb.WithContext(c)
	if req.ID != "" {
		sql = sql.Where(npcTaskDb.ID.Eq(req.ID))
	}
	if req.NpcID != "" {
		sql = sql.Where(npcTaskDb.NpcID.Eq(req.NpcID))
	}
	total, err := sql.Count()
	if err != nil {
		log.Printf("npcTaskDb count failed, err:%v\n", err)
		return nil, util.BuildErrorInfo("npcTaskDb count failed, err:%v", err)
	}
	npcTaskList, err := sql.Offset((req.PageNo - 1) * req.PageSize).Limit(req.PageSize).Find()
	if err != nil {
		log.Printf("npcTaskDb query failed, err:%v\n", err)
		return nil, util.BuildErrorInfo("npcTaskDb query failed, err:%v\n", err)
	}
	res := &model_view.GetNpcTaskResp{
		Total: total,
		Data:  npcTaskList,
	}
	return res, nil
}
