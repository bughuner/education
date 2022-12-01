package task

import (
	"education/database"
	"education/model"
	"education/util"
	"github.com/gin-gonic/gin"
	"log"
)

func GetTaskById(c *gin.Context, taskIds []string) ([]*model.Task, error){
	taskDb := database.Query.Task
	taskList, err := taskDb.WithContext(c).Where(taskDb.ID.In(taskIds...)).Find()
	if err != nil {
		log.Printf("taskDb query failed, taskIds:%v, err:%v\n", taskIds, err)
		return nil, util.BuildErrorInfo("taskDb query failed, taskIds:%v, err:%v", taskIds, err)
	}
	return taskList, nil
}
