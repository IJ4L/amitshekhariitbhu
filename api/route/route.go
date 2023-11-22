package route

import (
	"time"

	"architecture.com/api/jwt"
	"architecture.com/bootstrap"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, gin *gin.Engine) {
	publicRouter := gin.Group("")

	// All Public APIs
	NewSignupRouter(env, timeout, db, publicRouter)
	NewLoginRouter(env, timeout, db, publicRouter)
	NewRefreshTokenRouter(env, timeout, db, publicRouter)

	protectedRouter := gin.Group("")
	
	// Middleware to verify AccessToken
	protectedRouter.Use(jwt.JwtAuthMiddleware(env.AccessTokenSecret))

	// All Private APIs
	NewTaskRouter(env, timeout, db, protectedRouter)
}
