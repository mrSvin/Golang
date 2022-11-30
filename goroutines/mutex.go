package main

import (
	"fmt"
	"sync"
	"time"
)

var counter = 0 //  общий ресурс

func Mutex() {

	var mutex sync.Mutex

	go counterFunc(0, &mutex)
	go counterFunc(1, &mutex)
	go counterFunc(2, &mutex)
	go counterFunc(3, &mutex)
	go counterFunc(4, &mutex)
	time.Sleep(3 * time.Second)

}

func counterFunc(number int, mutex *sync.Mutex) {
	mutex.Lock()
	counter = 0
	for k := 1; k <= 5; k++ {
		counter++
		fmt.Println("Goroutine", number, "-", counter)
	}
	mutex.Unlock()

}
