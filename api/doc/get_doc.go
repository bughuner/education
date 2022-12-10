package doc

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

func GetDocApi(c *gin.Context) {
	var req model_view.GetDocReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("ShouldBindJSON GetDocReq failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	if err = checkGetParam(&req); err != nil {
		log.Printf("checkParam failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.NoParams, err.Error())
		return
	}
	res, err := getDoc(c, &req)
	if err != nil {
		log.Printf("getDoc failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.OperationErr, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, res)
}

func checkGetParam(req *model_view.GetDocReq) error {
	if req.PageNo == 0 {
		req.PageNo = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 100
	}
	return nil
}

func getDoc(c *gin.Context, req *model_view.GetDocReq) (*model_view.GetDocResp, error) {
	docDb := database.Query.Doc
	sql := docDb.WithContext(c)
	if req.ID != "" {
		sql = sql.Where(docDb.ID.Eq(req.ID))
	}
	if req.Author != "" {
		sql = sql.Where(docDb.Author.Eq(req.Author))
	}
	total, err := sql.Count()
	if err != nil {
		log.Printf("docDb count failed, err:%v\n", err)
		return nil, util.BuildErrorInfo("docDb count failed, err:%v", err)
	}
	docList, err := sql.Offset(req.PageNo - 1).Limit(req.PageSize).Find()
	if err != nil {
		log.Printf("docDb query failed, err:%v\n", err)
		return nil, util.BuildErrorInfo("docDb query failed, err:%v\n", err)
	}
	getDocRes := make([]*model.Doc, len(docList))
	for i, item := range docList {
		getDocRes[i] = &model.Doc{
			ID:      item.ID,
			Link:    item.Link,
			Author:  item.Author,
			Content: item.Content,
			Status:  item.Status,
		}
	}
	res := &model_view.GetDocResp{
		Total: total,
		Data:  getDocRes,
	}
	return res, nil
}
