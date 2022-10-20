package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/husfuu/go-gram/auth"
	"github.com/husfuu/go-gram/helper"
)

func Authorization(c *gin.Context) {
	authorizationHeader := c.GetHeader("Authorization")
	if !strings.Contains(authorizationHeader, "Bearer") {
		c.JSON(http.StatusUnauthorized, helper.NewResponse(http.StatusUnauthorized, nil, errors.New("the request is allowed for logged in")))
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	bearerToken := strings.Replace(authorizationHeader, "Bearer ", "", -1)

	id, err := auth.ParseToken(bearerToken)

	if err != nil {
		c.JSON(http.StatusUnauthorized, helper.NewResponse(http.StatusUnauthorized, nil, err))
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	c.Set("user_id", id)
	c.Next()
}
