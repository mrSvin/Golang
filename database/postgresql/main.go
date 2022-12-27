package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type login struct {
	name     string
	password string
}

func main() {
	db := connectToDb()
	defer db.Close()
	//writeData(db)
	selectData(db)

}

func connectToDb() *sql.DB {
	connStr := "user=postgres password=root dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	return db
}

func writeData(db *sql.DB) {
	_, err := db.Exec("INSERT INTO auth.auth_jwt (name, password) VALUES ($1, $2)", "admin", "test")
	if err != nil {
		panic(err)
	}
	fmt.Println("Данные добавлены") // не поддерживается
}

func selectData(db *sql.DB) {
	rows, err := db.Query("select * from auth.auth_jwt")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	logins := []login{}

	for rows.Next() {
		p := login{}
		err := rows.Scan(&p.name, &p.password)
		if err != nil {
			fmt.Println(err)
			continue
		}
		logins = append(logins, p)
	}
	for _, p := range logins {
		fmt.Println(p.name, p.password)
	}
}
