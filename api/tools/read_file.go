package tools

import (
	"education/common"
	errno "education/common/erron"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"strings"
)

func ReadFileApi(c *gin.Context) {
	name := c.Query("file_name")
	isExist, err := pathExists(name)
	if err != nil {
		log.Printf("pathExists failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	if !isExist {
		common.SendResponse(c, errno.ErrParams, "文件不存在")
		return
	}
	// 获取文件名称
	str := strings.Split(name, "/")
	fileName := str[len(str)-1]
	data, _ := ioutil.ReadFile(name)
	c.Header("Content-Type", "application/octet-stream")
	disposition := fmt.Sprintf("attachment; filename=%v", fileName)
	c.Header("Content-Disposition", disposition)
	c.Header("Content-Transfer-Encoding", "binary")
	c.Data(200, "application/octet-stream", data)
}
