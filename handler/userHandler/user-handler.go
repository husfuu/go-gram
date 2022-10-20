package userHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/husfuu/go-gram/dto"
	"github.com/husfuu/go-gram/helper"
	service "github.com/husfuu/go-gram/service/userService"
)

type UserHandler interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type handler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) UserHandler {
	return &handler{service: service}
}

func (h handler) Register(ctx *gin.Context) {
	input := new(dto.RequestRegister)

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, helper.NewResponse(http.StatusUnprocessableEntity, nil, err))
		return
	}

	response, err := h.service.RegisterUser(*input)

	if err != nil {
		ctx.JSON(
			helper.GetErrorStatusCode(err),
			helper.NewResponse(
				helper.GetErrorStatusCode(err),
				nil,
				err,
			),
		)
		return
	}

	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusCreated, response, err))
}

func (h handler) Login(ctx *gin.Context) {
	input := new(dto.RequestLogin)
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, helper.NewResponse(http.StatusUnprocessableEntity, nil, err))
		return
	}
	response, err := h.service.Login(*input)

	if err != nil {
		ctx.JSON(helper.GetErrorStatusCode(err), helper.NewResponse(helper.GetErrorStatusCode(err), nil, err))
		return
	}
	ctx.JSON(http.StatusOK, helper.Response{Status: http.StatusOK, Data: response, Error: err})
}

func (h handler) Update(ctx *gin.Context) {
	input := new(dto.RequestRegister)

	if err := ctx.ShouldBind(input); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, helper.NewResponse(http.StatusUnprocessableEntity, nil, err))
		return
	}
	id := ctx.MustGet("user_id")
	input.ID = id.(string)
	response, err := h.service.Update(*input)

	if err != nil {
		ctx.JSON(helper.GetErrorStatusCode(err), helper.NewResponse(helper.GetErrorStatusCode(err), nil, err))
		return
	}
	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, response, nil))
}

func (h handler) Delete(ctx *gin.Context) {
	id := ctx.MustGet("user_id")

	err := h.service.DeleteByID(id.(string))

	if err != nil {
		ctx.JSON(helper.GetErrorStatusCode(err), helper.NewResponse(helper.GetErrorStatusCode(err), nil, err))
		return
	}

	message := map[string]interface{}{"message": "your account has been successfully deleted"}
	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, message, nil))
}
