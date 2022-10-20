package socialmediaHandler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/husfuu/go-gram/dto"
	"github.com/husfuu/go-gram/helper"
	"github.com/husfuu/go-gram/service/socialmediaService"
)

type SocialMediaHandler interface {
	Create(ctx *gin.Context)
	GetSocialMedias(ctx *gin.Context)
	Update(ctx *gin.Context)
	DeleteByID(ctx *gin.Context)
}

type handler struct {
	service socialmediaService.SocialMediaService
}

func NewSocialMediaHandler(service socialmediaService.SocialMediaService) SocialMediaHandler {
	return &handler{service: service}
}

func (h handler) Create(ctx *gin.Context) {
	input := new(dto.RequestSocialMedia)

	err := ctx.ShouldBind(input)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, helper.NewResponse(http.StatusUnprocessableEntity, nil, err))
		return
	}
	input.UserID = ctx.MustGet("user_id").(string)

	response, err := h.service.Create(*input)
	if err != nil {
		ctx.JSON(helper.GetErrorStatusCode(err), helper.NewResponse(helper.GetErrorStatusCode(err), nil, err))
	}

	ctx.JSON(http.StatusCreated, helper.NewResponse(http.StatusCreated, response, nil))
}

func (h handler) GetSocialMedias(ctx *gin.Context) {
	response, err := h.service.GetSocialMedias()
	fmt.Println("ini response yg di dapet di handler", response)
	if err != nil {
		ctx.JSON(helper.GetErrorStatusCode(err), helper.NewResponse(helper.GetErrorStatusCode(err), nil, err))
		return
	}
	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, response, nil))
}

func (h handler) Update(ctx *gin.Context) {
	input := new(dto.RequestSocialMedia)

	err := ctx.ShouldBind(input)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, helper.NewResponse(http.StatusUnprocessableEntity, nil, err))
		return
	}
	socialMediaID := ctx.Param("social_media_id")
	input.UserID = ctx.MustGet("user_id").(string)

	input.ID = socialMediaID

	response, err := h.service.Update(*input)
	if err != nil {
		ctx.JSON(helper.GetErrorStatusCode(err), helper.NewResponse(helper.GetErrorStatusCode(err), nil, err))
	}

	ctx.JSON(http.StatusCreated, helper.NewResponse(http.StatusCreated, response, nil))
}

func (h handler) DeleteByID(ctx *gin.Context) {
	socialMediaID := ctx.Param("social_media_id")

	err := h.service.Delete(socialMediaID)

	if err != nil {
		ctx.JSON(helper.GetErrorStatusCode(err), helper.NewResponse(helper.GetErrorStatusCode(err), nil, err))
		return
	}
	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, map[string]interface{}{"message": "Your social media has been successfully deleted"}, nil))
}
