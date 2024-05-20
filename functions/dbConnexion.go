package dbConnexion

import (
	"database/sql"
	"fmt"
	"log"

	localSettings "crm/gogorigo/localSettings"

	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	psqlInfo := localSettings.DBPath

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected!")

	return db
}
