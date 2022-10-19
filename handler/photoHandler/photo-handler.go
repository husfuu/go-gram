package photoHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/husfuu/go-gram/dto"
	"github.com/husfuu/go-gram/helper"
	"github.com/husfuu/go-gram/service/photoService"
)

type PhotoHandler interface {
	Create(ctx *gin.Context)
	GetPhotos(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type handler struct {
	service photoService.PhotoService
}

func NewPhotoHandler(service photoService.PhotoService) PhotoHandler {
	return &handler{service: service}
}

func (h handler) Create(ctx *gin.Context) {
	input := new(dto.RequestPhoto)
	err := ctx.ShouldBind(input)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.NewResponse(http.StatusUnprocessableEntity, nil, err))
		return
	}
	userID := ctx.MustGet("user_id")
	input.UserID = userID.(string)
	response, err := h.service.Create(*input)
	if err != nil {
		ctx.JSON(helper.GetErrorStatusCode(err), helper.NewResponse(helper.GetErrorStatusCode(err), nil, err))
		return
	}
	ctx.JSON(http.StatusCreated, helper.NewResponse(http.StatusCreated, response, nil))
}

func (h handler) GetPhotos(ctx *gin.Context) {
	response, err := h.service.GetPhotos()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.NewResponse(http.StatusBadRequest, nil, err))
	}
	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, response, nil))
}

func (h handler) Update(ctx *gin.Context) {
	input := new(dto.RequestPhoto)
	err := ctx.ShouldBind(input)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.NewResponse(http.StatusUnprocessableEntity, nil, err))
		return
	}
	photoParamID := ctx.Param("photo_id")
	userID := ctx.MustGet("user_id")
	input.UserID = userID.(string)

	response, err := h.service.Update(*input, photoParamID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.NewResponse(helper.GetErrorStatusCode(err), nil, err))
		return
	}
	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, response, nil))
}

func (h handler) Delete(ctx *gin.Context) {
	photoParamID := ctx.Param("photo_id")

	err := h.service.Delete(photoParamID)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, helper.NewResponse(helper.GetErrorStatusCode(err), nil, err))
		return
	}
	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, "Your Photo has been successfully deleted", nil))
}
