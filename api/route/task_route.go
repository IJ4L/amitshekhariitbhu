package route

import (
	"time"

	"architecture.com/api/controller"
	"architecture.com/bootstrap"
	"architecture.com/repository"
	"architecture.com/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewTaskRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	tr := repository.NewTaskRepository(db)
	tc := controller.TaskController{
		TaskUsecase: usecase.NewTaskUsecase(tr, timeout),
	}
	group.GET("/task", tc.FetchByUserID)
	group.POST("/task", tc.Create)
}
