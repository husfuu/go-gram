package main

import (
	"log"

	"github.com/husfuu/go-gram/server"
	_ "github.com/joho/godotenv/autoload"
)

func main() {

	err := server.Start()
	if err != nil {
		log.Fatalln(err)
	}

}
