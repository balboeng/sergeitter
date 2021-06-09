package main

import (
	"log"

	"github.com/balboeng/sergeitter/db"
	"github.com/balboeng/sergeitter/handlers"
)

func main() {
	if db.CheckConnection() == 0 {
		log.Fatal("WITHOUT DB CONNECTION")
		return
	}
	handlers.Handlers()
}
