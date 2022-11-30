package main

import (
	"fmt"
	"time"
)

func Synchronized() {
	intCh := make(chan int)
	go solutionStepFirst(10, intCh)
	solutionStepTwo(<-intCh)
}

func solutionStepFirst(n int, chanIn chan int) {
	result := n * n
	chanIn <- result
}

func solutionStepTwo(chanIN int) {
	result := chanIN - 30
	fmt.Println(result)
}

func SynchronizedTwoWhile() {
	intCh := make(chan int, 2)
	intCh <- 1
	go whileFirst(intCh)
	go whileTwo(intCh)
	time.Sleep(5 * time.Second)
}

func whileFirst(chanIn chan int) {
	for i := 0; i < 10; i++ {
		if <-chanIn == 1 {
			fmt.Println("first: ", i)
			time.Sleep(30 * time.Millisecond)
			chanIn <- 0
		}
	}
}

func whileTwo(chanIn chan int) {
	for i := 0; i < 10; i++ {
		if <-chanIn == 0 {
			fmt.Println("two: ", i)
			time.Sleep(10 * time.Millisecond)
			chanIn <- 1
		}
	}
}
