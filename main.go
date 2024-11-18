package main

import (
	"log"

	"github.com/GDG-KHU-Side/backend-side-project/config"
	"github.com/GDG-KHU-Side/backend-side-project/db"
)

func main() {
	conf := config.GetDBConfig()

	err := db.InitDB(conf)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.DB.Close()

}
