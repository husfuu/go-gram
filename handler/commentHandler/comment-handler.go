package commentHandler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/husfuu/go-gram/dto"
	"github.com/husfuu/go-gram/helper"
	commentservice "github.com/husfuu/go-gram/service/commentService"
)

type CommentHandler interface {
	Create(ctx *gin.Context)
	GetComments(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type handler struct {
	service commentservice.CommentService
}

func NewCommentHandler(service commentservice.CommentService) CommentHandler {
	return &handler{service: service}
}

func (h handler) Create(ctx *gin.Context) {
	input := new(dto.RequestComment)
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

func (h handler) GetComments(ctx *gin.Context) {
	responses, err := h.service.GetComments()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.NewResponse(helper.GetErrorStatusCode(err), nil, err))
		return
	}
	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, responses, nil))
}

func (h handler) Update(ctx *gin.Context) {
	commentParamID := ctx.Param("comment_id")
	fmt.Println(commentParamID)
	input := new(dto.RequestCommentUpdate)

	userID := ctx.MustGet("user_id")
	input.UserID = userID.(string)

	err := ctx.ShouldBind(input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.NewResponse(http.StatusBadRequest, nil, err))
		return
	}
	updatedComment, err := h.service.Update(*input, commentParamID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.NewResponse(http.StatusBadRequest, nil, err))
		return
	}

	ctx.JSON(http.StatusAccepted, helper.NewResponse(http.StatusAccepted, updatedComment, nil))
}

func (h handler) Delete(ctx *gin.Context) {
	commentParamID := ctx.Param("comment_id")

	err := h.service.Delete(commentParamID)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, helper.NewResponse(helper.GetErrorStatusCode(err), nil, err))
		return
	}
	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, "your comment has been successfully deleted", nil))
}
