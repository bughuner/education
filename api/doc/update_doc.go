package doc

import (
	"education/common"
	errno "education/common/erron"
	"education/database"
	"education/model"
	"education/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

func UpdateDocApi(c *gin.Context) {
	var req model.Doc
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("ShouldBindJSON doc failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	if err = checkUpdateParam(req.ID); err != nil {
		log.Printf("checkParam failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.NoParams, err.Error())
		return
	}
	doc, err := updateDoc(c, req.ID, req.Link, req.Author, req.Content)
	if err != nil {
		log.Printf("addDoc failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.InternalServerError, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, doc)
}

func checkUpdateParam(id string) error {
	if id == "" {
		return util.BuildErrorInfo("id 为空")
	}
	return nil
}

func updateDoc(c *gin.Context, id, link, author, content string) (*model.Doc, error) {
	docDb := database.Query.Doc
	doc, err := docDb.WithContext(c).Where(docDb.ID.Eq(id)).First()
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Printf("docDb query failed, err:%v", err)
		return nil, err
	}
	if doc == nil {
		return nil, util.BuildErrorInfo("未找到文章")
	}
	if link != "" {
		doc.Link = link
	}
	if author != "" {
		doc.Author = author
	}
	if content != "" {
		doc.Content = content
	}
	err = docDb.WithContext(c).Save(doc)
	if err != nil {
		log.Printf("docDb save failed, err:%v", err)
		return nil, err
	}
	return doc, nil
}
