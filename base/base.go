package main

import "fmt"

func main() {
	switchInput(8)
	fmt.Println(returnIndexArray(1))
	ifElse(7, 3)
	whileFor()
	continueFor()
	breakFor()

	randomCountArgsSum(1, 2, 3, 4)
	randomCountArgsSum(5, 6, 7, 2, 3)
	var nums = []int{5, 6, 7, 2, 3}
	randomCountArgsSum(nums...)
	randomCountArgsSum([]int{5, 6, 7, 2, 3}...)

	fmt.Println(returnManyVals())

	f := returnFunc(2)
	fmt.Println(f(3, 4))

	deferPanic()

	anonimFunc := anonimFuncInput()
	fmt.Println(anonimFunc(3, 3))

	dynamicArraySlice()

	mapOperations()

}

func switchInput(a int) {
	switch a {
	case 9:
		fmt.Println("a = 9")
	case 8:
		fmt.Println("a = 8")
	case 7:
		fmt.Println("a = 7")
	default:
		fmt.Println("значение переменной a не определено")
	}
}

func returnIndexArray(index int) int {
	var numbers = [5]int{1, 2, 3, 4, 5}
	return numbers[index]
}

func ifElse(number1 int, number2 int) {
	if number1 < number2 {
		fmt.Println(number1, " меньше ", number2)
	} else {
		fmt.Println(number1, " больше ", number2)
	}
}

func whileFor() {
	var users = [3]string{"Tom", "Alice", "Kate"}

	for i := 0; i < len(users); i++ {
		fmt.Println(users[i])
	}
	for index, value := range users {
		fmt.Println(index, value)
	}
}

func continueFor() {
	var numbers = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var sum = 0

	for _, value := range numbers {
		if value > 5 {
			continue // переходим к следующей итерации
		}
		sum += value
		fmt.Println("value:", value)
	}
	fmt.Println("Sum:", sum)

}

func breakFor() {
	var numbers = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var sum = 0

	for _, value := range numbers {
		if value > 4 {
			break // если число больше 4 выходим из цикла
		}
		sum += value
	}
	fmt.Println("Sum:", sum)

}

func randomCountArgsSum(numbers ...int) {
	var sum = 0
	for _, number := range numbers {
		sum += number
	}
	fmt.Println("sum = ", sum)
}

func returnManyVals() (int, string, bool, float32) {
	return 6, "sd", true, 3.2
}

func add(x int, y int) int      { return x + y }
func subtract(x int, y int) int { return x - y }
func multiply(x int, y int) int { return x * y }
func returnFunc(n int) func(int, int) int {
	if n == 1 {
		return add
	} else if n == 2 {
		return subtract
	} else {
		return multiply
	}
}

func anonimFuncInput() func(int, int) int {
	return func(x int, y int) int { return x + y }
}

func deferPanic() {
	defer fmt.Println("end")
	fmt.Println("start")
	//panic("Stop!")
	fmt.Println("work")
}

func dynamicArraySlice() {

	var users []string
	var admins = []string{"Alex", "Vasya", "Tom", "Max", "Jerry"}

	users = append(users, admins[1])
	users = append(users, "Olya")

	for _, value := range admins {
		fmt.Println(value)
	}
	fmt.Println()
	for _, value := range users {
		fmt.Println(value)
	}
	fmt.Println()
	//delete 3 element 3
	var n = 2
	admins = append(admins[:n], admins[n+1:]...)

	for _, value := range admins {
		fmt.Println(value)
	}

}

func mapOperations() {
	var people = map[string]int{"Tom": 1, "Bob": 2, "Alex": 30}
	people["Sava"] = 4
	delete(people, "Tom")
	fmt.Println(people)
	fmt.Println(people["Alex"])

}
