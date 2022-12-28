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

func ReadPassword(loginInput string) string {
	db := connectToDb()
	defer db.Close()
	return findPassword(db, loginInput)
}

func connectToDb() *sql.DB {
	connStr := "user=postgres password=root dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	return db
}

func findPassword(db *sql.DB, loginInput string) string {
	rows, err := db.Query("select * from auth.auth_jwt where name = $1", loginInput)
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
		//fmt.Println(p.name, p.password)
		return p.password
	}
	return ""
}
