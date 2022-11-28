package main

import "fmt"

func main() {

	simplePoint()

	newPoint()

	y := 3
	noChangeInputVar(y)
	fmt.Println(y)

	changeInputVar(&y)
	fmt.Println(y)

	p1 := resultPointer(7)
	p2 := resultPointer(10)
	p3 := resultPointer(14)
	fmt.Println(*p1)
	fmt.Println(*p2)
	fmt.Println(*p3)

}

func simplePoint() {
	var x int = 10             // определяем переменную
	var p *int                 // определяем указатель
	p = &x                     // указатель получает адрес переменной
	fmt.Println("Address:", p) // значение указателя - адрес переменной x
	fmt.Println("Value:", *p)  // значение переменной x

	*p = 25
	fmt.Println("Address:", p)   // значение указателя - адрес переменной x
	fmt.Println("Value *p:", *p) // значение переменной x
	fmt.Println("Value x:", x)   // значение переменной x

}

func newPoint() {
	p := new(int)
	*p = 20
	fmt.Println(*p)
	fmt.Println(p)
}

func noChangeInputVar(x int) {
	x = x * x
}

func changeInputVar(x *int) {
	*x = (*x) * (*x)
}

func resultPointer(x int) *int {
	p := new(int)
	*p = x * x
	return p
}
