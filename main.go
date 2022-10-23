package main

import (
	"log"

	"github.com/husfuu/go-gram/server"
	_ "github.com/joho/godotenv/autoload"
)

// @title GoGram API
// @version 1.0
// @description This is API for completing hacktiv8 final project
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /

func main() {
	// gin.SetMode(gin.ReleaseMode)
	err := server.Start()
	if err != nil {
		log.Fatalln(err)
	}

}
