package config

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-restapi-service/controllers"
	"log"
	"os"
)

func Connect() *pg.DB {
	opts := &pg.Options{
		User:     "postgres",
		Password: "postgres",
		Addr:     "localhost:5432",
		Database: "postgres",
	}

	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Printf("Failed to connect")
		os.Exit(100)
	}

	log.Printf("Connected to db")

	controllers.CreateGroupTable(db)
	controllers.CreateUserTable(db)
	controllers.InitializeDB(db)

	return db
}
