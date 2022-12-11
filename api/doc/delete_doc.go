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

func DeleteDocApi(c *gin.Context) {
	var req model_view.DeleteDocReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("ShouldBindJSON DeleteDocReq failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	err = deleteDoc(c, req.DocIds)
	if err != nil {
		log.Printf("deleteDoc failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.OperationErr, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, nil)
}

func deleteDoc(c *gin.Context, ids []string) error {
	docDb := database.Query.Doc
	_, err := docDb.WithContext(c).Where(docDb.ID.In(ids...)).Delete(&model.Doc{})
	if err != nil {
		log.Printf("docDb delete failed, err:%v", err)
		return util.BuildErrorInfo("docDb delete failed, err:%v", err)
	}
	return nil
}
