package models

import (
	"api/db"
	"api/structs"
	"log"
)

func FindAllEndpoints() []structs.Endpoint {
	db := db.ConnectDB()
	defer db.Close()
	selectAll, err := db.Query("SELECT * FROM Endpoints ORDER BY id ASC")
	if err != nil {
		log.Println(err.Error())
	}

	endpoints := []structs.Endpoint{}

	for selectAll.Next() {
		e := structs.Endpoint{}
		err = selectAll.Scan(&e.Id, &e.Url, &e.Endpoint, &e.Alertname)
		if err != nil {
			log.Println(err.Error())
		}
		endpoints = append(endpoints, e)
	}
	return endpoints
}

func findEndpoint(url, endpoint string) structs.Endpoint {
	db := db.ConnectDB()
	defer db.Close()
	var e structs.Endpoint
	db.QueryRow("SELECT * FROM Endpoints WHERE url=? AND endpoint=?", url, endpoint).Scan(&e.Id, &e.Url, &e.Endpoint)
	return e
}

func InsertEndpoint(url, endpoint, alertname string) bool {
	db := db.ConnectDB()
	defer db.Close()
	e := findEndpoint(url, endpoint)
	if (structs.Endpoint{}) != e {
		return false
	} else {

		insertSQL := `INSERT INTO Endpoints(url, endpoint, alertname) VALUES (?, ?, ?)`
		stm, err := db.Prepare(insertSQL)
		if err != nil {
			log.Println("Error when preparing the statement")
			return false
		}
		_, err = stm.Exec(url, endpoint, alertname)
		if err != nil {
			log.Println("Error when executing the statement")
			return false
		}

	}
	return true
}

func DeleteEndpoint(id string) bool {
	db := db.ConnectDB()
	defer db.Close()
	deleteSQL := `DELETE FROM Endpoints WHERE Id=?`
	stm, err := db.Prepare(deleteSQL)

	if err != nil {
		log.Println("Erro ao executar statement")
		return false
	}

	_, err = stm.Exec(id)
	if err != nil {
		log.Println("Erro ao executar statement")
		return false
	}
	return true
}
