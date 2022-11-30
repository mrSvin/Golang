package main

import (
	"fmt"
)

func main() {
	//go factorial(100)
	//go factorial(80)
	//go factorial(120)
	//time.Sleep(5 * time.Second)
	//NoBufferChan()
	//BufferChan()
	//OnlyReadOnlyWriteChannels()
	//ReturnChannel()
	//CloseChannel()
	//Synchronized()
	//SynchronizedTwoWhile()
	//SendChannel()
	//Mutex()
	WaitGroup()

}

func factorial(n int) {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	fmt.Println(result, " - n: ", n)

}
