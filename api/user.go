package api

import (
	"education/common"
	errno "education/common/erron"
	"education/database"
	"education/model"
	_ "education/model"
	"education/util"
	"github.com/gin-gonic/gin"
	"log"
)

// 新建工单
func GetUserApi(c *gin.Context) {

}

func AddUserApi(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Printf("ShouldBindJSON user failed,err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	user.ID = util.GetUUID()

	userDb := database.Query.User
	err = userDb.WithContext(c).Create(&user)
	if err != nil {
		log.Printf("userDb create user failed,err:%v\n", err)
		common.SendResponse(c, errno.AddErr, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, user)
}
