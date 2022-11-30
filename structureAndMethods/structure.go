package main

import "fmt"

type person struct {
	name string
	age  int
}

type contact struct {
	email string
	phone string
}

type nestedContactPersone struct {
	personeInfo person
	contactInfo contact
}

func main() {
	createChangePersone("Alex")

	sava := copyStructure("Sava")
	fmt.Println(sava.age)
	fmt.Println(sava.name)

	var archi = person{"Archi", 23}
	createNestedPersone(archi)

	archi.methodUpdateAge(24)
	fmt.Println(archi)

}

func createChangePersone(name string) {
	var newPersone = person{name, 23}
	fmt.Println(newPersone.name)
	fmt.Println(newPersone.age)
	newPersone.age = 24
	fmt.Println(newPersone.age)
}

func copyStructure(name string) *person {
	return &person{
		name: name,
		age:  4,
	}
}

func createNestedPersone(personInput person) {
	var archi = nestedContactPersone{
		personeInfo: personInput,
		contactInfo: contact{
			email: "tom@gmail.com",
			phone: "+1234567899",
		},
	}
	fmt.Println(archi.personeInfo)
	fmt.Println(archi.contactInfo)
}

func (p *person) methodUpdateAge(newAge int) {
	p.age = newAge
}
