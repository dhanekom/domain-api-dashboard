package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gitlab.host-h.net/skunkworks/domain-api-dashboard/database"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DB_CONN_STRING")
	db, err := database.ConnectToDB(dsn)
	if err != nil {
		log.Fatalf("unable to connect to db: %s", err)
	}
	defer db.Close()

	dbRepo := database.NewMysqlDBRepo(db)

	handlerRepo := newHandlerRepo(dbRepo)

	router := routes(handlerRepo)

	router.Logger.Fatal(router.Start("127.0.0.1:42069"))
}
