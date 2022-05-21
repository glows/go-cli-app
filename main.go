package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/dixonwille/wmenu"
)

func main() {

	//connec to database
	db, err := sql.Open("sqlite3", "./people.db")
	checkErr(err)

	// defer close
	defer db.Close()

	menu := wmenu.NewMenu("what would you like to do?")

	menu.Action(func(o []wmenu.Opt) error { handleFunc(db, o); return nil })

	menu.Option("Add a new Person", 0, true, nil)
	menu.Option("Find a Person", 1, false, nil)
	menu.Option("Update a Person's informaion", 2, false, nil)
	menu.Option("Delete a Person by ID", 3, false, nil)
	menuerr := menu.Run()
	if menuerr != nil {
		log.Fatal(menuerr)
	}

}

func handleFunc(db *sql.DB, opts []wmenu.Opt) {
	switch opts[0].Value {

	case 0:
		fmt.Println("Adding a new Person")
		fmt.Print("Enter a first name")
		firstName, _ := reader.ReadString('\n')
		fmt.Print("Enter a last name")
		lastName, _ := reader.ReadString('\n')
		fmt.Print("Enter an email address:")
		email, _ := reader.ReadString('\n')
		fmt.Print("Enter an IP address:")
		ipAddress, _ := reader.ReadString('\n')

		newPerson := person{
			first_name: firstName,
			last_name:  lastName,
			email:      email,
			ip_address: ipAddress,
		}

		addPerson(db, newPerson)

		break

	case 1:
		fmt.Println("Finding a Persion")
	case 2:
		fmt.Println("Update a Persion's information")
	case 3:
		fmt.Println("Deleting a Persion by ID")
	case 4:
		fmt.Println("Quitting application")
	}
}
