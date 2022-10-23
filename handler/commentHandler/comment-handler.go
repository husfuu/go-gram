package commentHandler

import (
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

// Create a comment
// @Tags comments
// @Summary Create a comment
// @Description Create a comment
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer + user token"
// @Param data body dto.RequestComment true "data"
// @Success 201 {object} helper.Response{data=dto.ResponseCreateComment} "CREATED"
// @Failure 400 {object} helper.Response{errors=helper.ExampleErrorResponse} "Bad Request"
// @Failure 401 {object} helper.Response{errors=helper.ExampleErrorResponse} "Unauthorization"
// @Router /comments [POST]
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

// Get comments
// @Tags comments
// @Summary Get comments
// @Description Get comments
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer + user token"
// @Success 200 {object} helper.Response{data=[]dto.ResponseGetComment} "OK"
// @Failure 400 {object} helper.Response{errors=helper.ExampleErrorResponse} "Bad Request"
// @Failure 401 {object} helper.Response{errors=helper.ExampleErrorResponse} "Unauthorization"
// @Router /comments [GET]
func (h handler) GetComments(ctx *gin.Context) {
	responses, err := h.service.GetComments()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.NewResponse(helper.GetErrorStatusCode(err), nil, err))
		return
	}
	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, responses, nil))
}

// Update a comment
// @Tags comments
// @Summary Update a comment
// @Description Update a comment
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer + user token"
// @Param commentID path int true "ID of the comment"
// @Param data body dto.RequestCommentUpdate true "data"
// @Success 200 {object} helper.Response{data=dto.ResponseCreateComment} "OK"
// @Failure 400 {object} helper.Response{errors=helper.ExampleErrorResponse} "Bad Request"
// @Failure 401 {object} helper.Response{errors=helper.ExampleErrorResponse} "Unauthorization"
// @Router /comments/:commentID [PUT]
func (h handler) Update(ctx *gin.Context) {
	commentParamID := ctx.Param("comment_id")

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

// Delete a comment
// @Tags comments
// @Summary Delete a comment
// @Description Delete a comment
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer + user token"
// @Param commentID path int true "ID of the comment"
// @Success 200 {object} helper.Response "OK"
// @Failure 400 {object} helper.Response{errors=helper.ExampleErrorResponse} "Bad Request"
// @Failure 401 {object} helper.Response{errors=helper.ExampleErrorResponse} "Unauthorization"
// @Router /comments/:commentID [DELETE]
func (h handler) Delete(ctx *gin.Context) {
	commentParamID := ctx.Param("comment_id")

	err := h.service.Delete(commentParamID)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, helper.NewResponse(helper.GetErrorStatusCode(err), nil, err))
		return
	}
	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, "your comment has been successfully deleted", nil))
}
