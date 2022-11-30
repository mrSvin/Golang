package main

import "fmt"

import "interface/manyInterface"

type Transport interface {
	move()
}

func drive(transport Transport) {
	transport.move()
}

type Car struct{}
type Aircraft struct{}

func (c Car) move() {
	fmt.Println("Автомобиль едет")
}
func (a Aircraft) move() {
	fmt.Println("Самолет летит")
}

func main() {

	var tesla = Car{}
	var boing = Aircraft{}
	tesla.move()
	boing.move()

	drive(tesla)
	drive(boing)

	StreamOperations()

	manyInterface.ManyInterface()
}
