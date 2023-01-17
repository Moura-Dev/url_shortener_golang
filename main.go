package main

import (
	"github.com/joho/godotenv"
	"url_shortner/db"
	"url_shortner/server"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	database, _ := db.StartDB()
	database.MustExec(db.Schema)

	defer database.Close()

	s := server.NewServer()

	s.Run()
}
