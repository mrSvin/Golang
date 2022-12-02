package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	createTable("names")
	insertData("names", "alex")
	insertData("names", "sava")
	insertData("names", "oksana")
	deleteData("alex")
	selectData("names")
}

func createTable(tableName string) {
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

func insertData(tableName string, nameData string) {
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
	id   int
	name string
}

func selectData(tableName string) {
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
		err := rows.Scan(&p.id, &p.name)
		if err != nil {
			fmt.Println(err)
			continue
		}
		names = append(names, p)
	}
	fmt.Println(names)

}

func deleteData(nameDelete string) {
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
