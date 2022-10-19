package server

import (
	"github.com/gin-gonic/gin"
	"github.com/husfuu/go-gram/handler/userHandler"
	"github.com/husfuu/go-gram/middleware"
	"github.com/husfuu/go-gram/repository/userRepository"
	"github.com/husfuu/go-gram/service/userService"
	"gorm.io/gorm"
)

func NewRouter(r *gin.Engine, db *gorm.DB) {
	userRepo := userRepository.NewUserRepository(db)
	userSrv := userService.NewUserService(userRepo)
	userHdlr := userHandler.NewUserHandler(userSrv)

	userRoute := r.Group("/users")
	userRoute.POST("/register", userHdlr.Register)
	userRoute.POST("/login", userHdlr.Login)
	userRoute.PUT("", middleware.Authorization, userHdlr.Update)
	userRoute.DELETE("", middleware.Authorization, userHdlr.Delete)
}
