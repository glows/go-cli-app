package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type person struct {
	id         int
	first_name string
	last_name  string
	email      string
	ip_address string
}

func addPerson(db *sql.DB, newPerson person) {
	stmt, _ := db.Prepare("INSERT INTO people (id, first_name, last_name, email,ip_address) VALUES (?,?,?,?,?)")

	stmt.Exec(nil, newPerson.first_name, newPerson.last_name, newPerson.email, newPerson.ip_address)
	defer stmt.Close()

	fmt.Printf("Added %v %v \n", newPerson.first_name, newPerson.last_name)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func findPersonById(db *sql.DB, ourID string) person {
	rows, _ := db.Query("SELECT id, first_name, last_name, email, ip_sddress FROM people WHERE id = '" + ourID + "'")
	defer rows.Close()

	ourPerson := person{}

	for rows.Next() {
		rows.Scan(&ourPerson.id, &ourPerson.first_name, &ourPerson.last_name, &ourPerson.email, &ourPerson.ip_address)

	}

	return ourPerson
}
