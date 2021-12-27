package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectDB() *sql.DB {
	db, err := sql.Open("sqlite3", "/var/lib/sql/database.db")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func CreateDB() {
	if _, err := os.Stat("/var/lib/sql/database.db"); os.IsNotExist(err) {

		log.Println("Creating database.db...")

		file, err := os.Create("/var/lib/sql/database.db")
		if err != nil {
			log.Println("database already present")
		}

		file.Close()
		log.Println("database.db created")
	}
}

func CreateTableHosts() {
	db := ConnectDB()
	defer db.Close()

	createHostTableSQL := `CREATE TABLE IF NOT EXISTS Endpoints  (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"url" TEXT,
		"endpoint" TEXT,
		"alertname" TEXT
	  );`

	log.Println("Creating table...")
	stm, err := db.Prepare(createHostTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	stm.Exec()
	log.Println("database up")
}
