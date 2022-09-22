package main

import (
	"asciiart/internal/server"
	"log"
)

func main() {
	if err := server.Server(); err != nil {
		log.Println(err.Error())
		return
	}
}
