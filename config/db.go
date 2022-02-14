package config

import (
	"context"
	"github.com/go-pg/pg/v10"
	"github.com/go-restapi-service/controllers"
	"log"
	"os"
	"time"
)

func Connect() *pg.DB {
	addr := "localhost:5432"
	if os.Getenv("ENV") == "docker" {
		addr = "host.docker.internal:5432"
	}
	opts := &pg.Options{
		User:     "postgres",
		Password: "postgres",
		Addr:     addr,
		Database: "postgres",
	}

	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Printf("Failed to connect")
		os.Exit(100)
	}

	log.Printf("Connected to db")

	//try 3 times, sometimes it takes longer for database to accept connections
	for i := 0; i < 3; i++ {
		if db.Ping(context.Background()) == nil {
			controllers.CreateGroupTable(db)
			controllers.CreateUserTable(db)
			controllers.InitializeDB(db)

			return db
		} else {
			time.Sleep(time.Second * 3)
		}
	}

	log.Printf("Connection failed: cannot ping db")
	os.Exit(100)
	return nil
}
