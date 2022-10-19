package server

import (
	"fmt"

	"github.com/husfuu/go-gram/config"
)

func Start() error {
	db, err := config.NewDbInit()

	if err != nil {
		panic(err)
	}

	// r := gin.Default()
	fmt.Println(db)
	return nil
}
