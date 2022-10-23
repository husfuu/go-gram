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

// Create a photo
// @Tags photos
// @Summary Create a new photo and store it in to database
// @Description Create a new photo
// @Accept json
// @Produce json
// @Param data body dto.RequestPhoto true "data"
// @Success 201 {object} helper.Response{data=dto.ResponseCreatePhoto} "CREATED"
// @Failure 400 {object} helper.Response{errors=helper.ExampleErrorResponse} "Bad Request"
// @Failure 401 {object} helper.Response{errors=helper.ExampleErrorResponse} "Unauthorization"
// @Router /photos [POST]
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

// GetPhotos a photo
// @Tags photos
// @Summary Get all photos
// @Description Get all photos
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer + user token"
// @Success 200 {object} helper.Response{data=[]dto.ResponseGetPhoto} "SUCCESS"
// @Failure 400 {object} helper.Response{errors=helper.ExampleErrorResponse} "Bad Request"
// @Failure 401 {object} helper.Response{errors=helper.ExampleErrorResponse} "Unauthorization"
// @Failure 404 {object} helper.Response{errors=helper.ExampleErrorResponse} "Not Found"
// @Router /photos [GET]
func (h handler) GetPhotos(ctx *gin.Context) {
	response, err := h.service.GetPhotos()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.NewResponse(http.StatusBadRequest, nil, err))
		return
	}
	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, response, nil))
}

// Update a photo
// @Tags photos
// @Summary Update a photo
// @Description Update a photo
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer + user token"
// @Param photoID path int true "ID of the photo"
// @Param data body dto.RequestPhoto true "data"
// @Success 200 {object} helper.Response{data=dto.ResponseUpdatePhoto} "SUCCESS"
// @Failure 400 {object} helper.Response{errors=helper.ExampleErrorResponse} "Bad Request"
// @Failure 401 {object} helper.Response{errors=helper.ExampleErrorResponse} "Unauthorization"
// @Router /photos/:photoID [PUT]
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

// Delete a photo
// @Tags photos
// @Summary Delete a photo
// @Description Delete a photo
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer + user token"
// @Param photoID path int true "ID of the photo"
// @Success 200 {object} helper.Response "SUCCESS"
// @Failure 400 {object} helper.Response{errors=helper.ExampleErrorResponse} "Bad Request"
// @Failure 401 {object} helper.Response{errors=helper.ExampleErrorResponse} "Unauthorization"
// @Router /photos/:photoID [DELETE]
func (h handler) Delete(ctx *gin.Context) {
	photoParamID := ctx.Param("photo_id")

	err := h.service.Delete(photoParamID)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, helper.NewResponse(helper.GetErrorStatusCode(err), nil, err))
		return
	}
	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, "Your Photo has been successfully deleted", nil))
}
