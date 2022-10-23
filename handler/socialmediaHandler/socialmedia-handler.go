package socialmediaHandler

import (
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

// Create new social media
// @Tags socialmedias
// @Summary Create new social media
// @Description Create social media
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer + user token"
// @Param data body dto.RequestSocialMedia true "data"
// @Success 201 {object} helper.Response{data=dto.ResponseCreateSocialMedia} "CREATED"
// @Failure 400 {object} helper.Response{errors=helper.ExampleErrorResponse} "Bad Request"
// @Router /socialmedias [POST]
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
		return
	}

	ctx.JSON(http.StatusCreated, helper.NewResponse(http.StatusCreated, response, nil))
}

// Get all social medias
// @Tags socialmedias
// @Summary Get all social medias
// @Description Get all social medias
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer + user token"
// @Success 200 {object} helper.Response{data=[]dto.ResponseGetSocialMedias} "SUCCESS"
// @Router /socialmedias [GET]
func (h handler) GetSocialMedias(ctx *gin.Context) {
	response, err := h.service.GetSocialMedias()

	if err != nil {
		ctx.JSON(helper.GetErrorStatusCode(err), helper.NewResponse(helper.GetErrorStatusCode(err), nil, err))
		return
	}
	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, response, nil))
}

// Update by id social media
// @Tags socialmedias
// @Summary Update by id social media
// @Description Update by id social media
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer + user token"
// @Param socialmediaid path int true "ID of the social media"
// @Param data body dto.RequestSocialMedia true "data"
// @Success 200 {object} helper.Response{data=dto.ResponseCreateSocialMedia} "SUCCESS"
// @Failure 400 {object} helper.Response{errors=helper.ExampleErrorResponse} "Bad Request"
// @Failure 404 {object} helper.Response{errors=helper.ExampleErrorResponse} "Record not found"
// @Router /socialmedias/:socialmediaid [PUT]
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

// Delete by id social media
// @Tags socialmedias
// @Summary Delete by id social media
// @Description Delete by id social media
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer + user token"
// @Param socialmediaid path int true "ID of the social media"
// @Success 200 {object} helper.Response "SUCCESS"
// @Failure 400 {object} helper.Response{errors=helper.ExampleErrorResponse} "Bad Request"
// @Failure 404 {object} helper.Response{errors=helper.ExampleErrorResponse} "Record not found"
// @Router /socialmedias/:socialmediaid [DELETE]
func (h handler) DeleteByID(ctx *gin.Context) {
	socialMediaID := ctx.Param("social_media_id")

	err := h.service.Delete(socialMediaID)

	if err != nil {
		ctx.JSON(helper.GetErrorStatusCode(err), helper.NewResponse(helper.GetErrorStatusCode(err), nil, err))
		return
	}
	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, map[string]interface{}{"message": "Your social media has been successfully deleted"}, nil))
}
