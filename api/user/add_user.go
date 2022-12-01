package user

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

// 添加用户
func AddUserApi(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Printf("ShouldBindJSON user failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	userInfo, err := addUser(c, &user)
	if err != nil {
		log.Printf("addUser failed, user:%v, err:%v\n", user, err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, userInfo)
}

func addUser(c *gin.Context, user *model.User) (*model.User, error){
	userDb := database.Query.User
	oldUser, err := userDb.WithContext(c).Where(userDb.Name.Eq(user.Name)).First()
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Printf("userDb query failed, err:%v\n", err)
		return nil, err
	}
	if oldUser != nil {
		log.Printf("用户名已经存在")
		return nil, util.BuildErrorInfo("用户名已经存在")
	}
	user.ID = util.GetUUID()
	err = userDb.WithContext(c).Create(user)
	if err != nil {
		log.Printf("userDb create user failed,err:%v\n", err)
		return nil, err
	}
	return user, nil
}