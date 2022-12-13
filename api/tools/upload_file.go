package tools

import (
	"education/common"
	errno "education/common/erron"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"mime/multipart"
	"os"
	"time"
)

func UploadFileApi(c *gin.Context) {
	f, err := c.FormFile("file")
	if err != nil {
		log.Printf("FormFile failed, err:%v\n", err)
		common.SendResponse(c, errno.InternalServerError, err.Error())
		return
	}
	// 获取当前绝对路径
	dst, err := os.Getwd()
	if err != nil {
		log.Printf("os.Getwd failed, err:%v\n", err)
		common.SendResponse(c, errno.InternalServerError, err.Error())
		return
	}
	// 创建存放地址
	path := fmt.Sprintf("%v/%v/", dst, "upload_image")
	isExist, err := pathExists(path)
	if err != nil {
		log.Printf("pathExists failed, err:%v\n", err)
		common.SendResponse(c, errno.InternalServerError, err.Error())
		return
	}
	// 文件夹不存在就直接创建
	if !isExist {
		os.Mkdir(path, 0644)
	}
	// 存储文件
	t := time.Now().Unix()
	fileDst := fmt.Sprintf("%v%v_%v", path, t, f.Filename)
	err = saveUploadedFile(c, f, fileDst)
	if err != nil {
		log.Printf("saveUploadedFile failed, err:%v\n", err)
		common.SendResponse(c, errno.InternalServerError, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, fileDst)
}

func saveUploadedFile(c *gin.Context, file *multipart.FileHeader, dst string) error {
	//打开请求发送的文件
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	//在本地创建一个文件
	out, err := os.OpenFile(dst, os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer out.Close()
	//把内容拷贝到本地文件
	_, err = io.Copy(out, src)
	return err
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
