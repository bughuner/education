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

func LoginApi(c *gin.Context) {
	var req model.User
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("ShouldBindJSON user failed,err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	if err = checkParam(req); err != nil {
		log.Printf("checkParam failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.NoParams, err.Error())
		return
	}
	user, err := login(c, req.Account, req.Password)
	if err != nil {
		log.Printf("login failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.OperationErr, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, user)
}

func checkParam(user model.User) error {
	if user.Account == "" {
		return util.BuildErrorInfo("账号不能为空")
	}
	if user.Password == "" {
		return util.BuildErrorInfo("密码不能为空")
	}
	return nil
}

func login(c *gin.Context, account, password string) (*model.User, error) {
	userDb := database.Query.User
	oldUser, err := userDb.WithContext(c).Where(userDb.Account.Eq(account)).First()
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Printf("userDb query failed, account:%v, password:%v, err:%v\n", account, password, err)
		return nil, util.BuildErrorInfo("userDb query failed, err:%v", err)
	}
	if oldUser == nil {
		return nil, util.BuildErrorInfo("用户不存在")
	}
	if oldUser.Password != password {
		return nil, util.BuildErrorInfo("密码错误")
	}
	return oldUser, nil
}

func getUserById(c *gin.Context, id string) (*model.User, error) {
	userDb := database.Query.User
	user, err := userDb.WithContext(c).Where(userDb.ID.Eq(id)).First()
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Printf("userDb query failed, id:%v, err:%v\n", id, err)
		return nil, util.BuildErrorInfo("userDb query failed, err:%v", err)
	}
	if user == nil {
		return nil, util.BuildErrorInfo("用户不存在")
	}
	return user, nil
}
