package main

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"

	dbConnexion "crm/gogorigo/functions"
)

var db *sql.DB

func main() {

	db = dbConnexion.Connect()
	defer db.Close()

	app := fiber.New()

	log.Fatal(app.Listen(":3000"))
}
