package server

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/husfuu/go-gram/config"
)

func Start() error {
	db, err := config.NewDbInit()

	if err != nil {
		panic(err)
	}

	r := gin.Default()
	NewRouter(r, db)

	r.Use(gin.Recovery())

	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = "8000"
	}

	r.Run(fmt.Sprintf(":%s", port))

	return nil
}
