package server

import (
	"github.com/gin-gonic/gin"
	"github.com/husfuu/go-gram/handler/commentHandler"
	"github.com/husfuu/go-gram/handler/photoHandler"
	"github.com/husfuu/go-gram/handler/socialmediaHandler"
	"github.com/husfuu/go-gram/handler/userHandler"
	"github.com/husfuu/go-gram/middleware"
	"github.com/husfuu/go-gram/repository/commentRepository"
	"github.com/husfuu/go-gram/repository/photoRepository"
	"github.com/husfuu/go-gram/repository/socialmediaRepository"
	"github.com/husfuu/go-gram/repository/userRepository"
	commentservice "github.com/husfuu/go-gram/service/commentService"
	"github.com/husfuu/go-gram/service/photoService"
	"github.com/husfuu/go-gram/service/socialmediaService"
	"github.com/husfuu/go-gram/service/userService"
	"gorm.io/gorm"
)

func NewRouter(r *gin.Engine, db *gorm.DB) {
	// user routes
	userRepo := userRepository.NewUserRepository(db)
	userSrv := userService.NewUserService(userRepo)
	userHdlr := userHandler.NewUserHandler(userSrv)

	userRoute := r.Group("/users")
	userRoute.POST("/register", userHdlr.Register)
	userRoute.POST("/login", userHdlr.Login)
	userRoute.PUT("", middleware.Authorization, userHdlr.Update)
	userRoute.DELETE("", middleware.Authorization, userHdlr.Delete)

	// photo routes
	photoRepo := photoRepository.NewPhotoRepository(db)
	photoSrv := photoService.NewPhotoService(photoRepo)
	photoHdlr := photoHandler.NewPhotoHandler(photoSrv)

	photoRoute := r.Group("/photos").Use(middleware.Authorization)
	photoRoute.POST("", photoHdlr.Create)
	photoRoute.GET("", photoHdlr.GetPhotos)
	photoRoute.PUT("/:photo_id", photoHdlr.Update)
	photoRoute.DELETE("/:photo_id", photoHdlr.Delete)

	// social media routes
	socialMediaRepo := socialmediaRepository.NewSocialMediaRepository(db)
	socialMediaSrv := socialmediaService.NewSocialMediaService(socialMediaRepo, photoRepo)
	socialMediaHdlr := socialmediaHandler.NewSocialMediaHandler(socialMediaSrv)

	socialMediaRoute := r.Group("/socialmedias").Use(middleware.Authorization)
	socialMediaRoute.POST("", socialMediaHdlr.Create)
	socialMediaRoute.GET("", socialMediaHdlr.GetSocialMedias)
	socialMediaRoute.PUT("/:social_media_id", socialMediaHdlr.Update)
	socialMediaRoute.DELETE("/:social_media_id", socialMediaHdlr.DeleteByID)

	// comment routes
	commentRepo := commentRepository.NewCommentRepository(db)
	commentSrv := commentservice.NewCommentService(commentRepo)
	commentHdlr := commentHandler.NewCommentHandler(commentSrv)

	commentRoute := r.Group("/comments").Use(middleware.Authorization)
	commentRoute.POST("", commentHdlr.Create)
	commentRoute.GET("", commentHdlr.GetComments)
	commentRoute.PUT("/:comment_id", commentHdlr.Update)
	commentRoute.DELETE("/:comment_id", commentHdlr.Delete)
}
