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

func NewProfileRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db)
	pc := controller.ProfileController{
		ProfileUsecase: usecase.NewProfileUsecase(ur, timeout),
	}
	group.GET("/profile", pc.Fetch)
}
