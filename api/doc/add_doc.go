package doc

import (
	"education/common"
	errno "education/common/erron"
	"education/database"
	"education/model"
	"education/util"
	"github.com/gin-gonic/gin"
	"log"
)

func AddDocApi(c *gin.Context) {
	var req model.Doc
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("ShouldBindJSON doc failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	if err = checkAddParam(&req); err != nil {
		log.Printf("checkParam failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.NoParams, err.Error())
		return
	}
	doc, err := addDoc(c, req.Link, req.Author, req.Content)
	if err != nil {
		log.Printf("addDoc failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.InternalServerError, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, doc)
}

func checkAddParam(doc *model.Doc) error {
	if doc.Link == "" {
		return util.BuildErrorInfo("链接为空")
	}
	if doc.Author == "" {
		return util.BuildErrorInfo("作者为空")
	}
	if doc.Content == "" {
		return util.BuildErrorInfo("内容为空")
	}
	return nil
}

func addDoc(c *gin.Context, link, auth, content string) (*model.Doc, error) {
	docDb := database.Query.Doc
	id := util.GetUUID()
	doc := &model.Doc{
		ID:      id,
		Link:    link,
		Author:  auth,
		Content: content,
	}
	err := docDb.WithContext(c).Create(doc)
	if err != nil {
		log.Printf("docDb create failed, err:%v", err)
		return nil, err
	}
	return doc, nil
}
