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

func NewSignupRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db)
	sc := controller.SignupController{
		SignupUsecase: usecase.NewSignUpUsecase(ur, timeout),
		Env:           env,
	}
	group.POST("/signup", sc.Signup)
}