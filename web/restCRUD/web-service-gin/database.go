package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func CreateTable(tableName string) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/golang")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	_, err = db.Exec("CREATE TABLE golang." + tableName + " (`id` INT NOT NULL AUTO_INCREMENT, `name` VARCHAR(45) NULL, PRIMARY KEY (`id`))")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("database created!")
}

func InsertData(tableName string, nameData string) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/golang")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	_, err = db.Exec("INSERT INTO golang."+tableName+" (`name`) VALUES (?)", nameData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(nameData + " добавлен!")
}

type name struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func SelectData(tableName string) []name {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/golang")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from " + tableName)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	names := []name{}
	for rows.Next() {
		p := name{}
		err := rows.Scan(&p.Id, &p.Name)
		if err != nil {
			fmt.Println(err)
			continue
		}
		names = append(names, p)
	}
	return names
}

func DeleteData(nameDelete string) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/golang")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("delete from names where name = (?)", nameDelete)
	if err != nil {
		panic(err)
	}
	fmt.Println(nameDelete + " удален!")

}
