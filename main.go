package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

// Person struct which will map to the record in Person table
type Person struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int    `json:"age"`
}

func (p *Person) queryAll(db *sql.DB) {
	results, err := db.Query("SELECT firstname, lastname, age FROM person")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var person Person

		// for each row, scan the result into our composite object
		err = results.Scan(&person.FirstName, &person.LastName, &person.Age)
		if err != nil {
			panic(err.Error())
		}

		log.Printf(person.FirstName + " " + person.LastName + " is " + strconv.Itoa(person.Age) + " years old")
	}
}

func (p *Person) query(db *sql.DB, id int) Person {
	var person Person

	row := db.QueryRow("SELECT id, firstname, lastname, age FROM person where id = ?", id)

	err := row.Scan(&person.ID, &person.FirstName, &person.LastName, &person.Age)
	if err != nil {
		panic(err.Error())
	}

	return person
}

func (p *Person) insert(db *sql.DB, firstname string, lastname string, age int) {
	insert, err := db.Query("INSERT INTO person ( firstname, lastname, age ) VALUES ( '" + firstname + "', '" + lastname + "', " + strconv.Itoa(age) + " )")

	if err != nil {
		panic(err.Error())
	}

	// be careful deferring queries if you are using transactions
	defer insert.Close()
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	p := &Person{}

	//p.insert(db, "Ana", "Watson", 21)

	p.queryAll(db)

	personRec := p.query(db, 1)
	fmt.Println(personRec)
}
