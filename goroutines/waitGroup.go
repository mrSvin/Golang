package main

import (
	"fmt"
	"sync"
	"time"
)

func WaitGroup() {
	var wg sync.WaitGroup
	wg.Add(3)
	go solutionForGroup(1, &wg)
	go solutionForGroup(2, &wg)
	go solutionForGroup(3, &wg)
	wg.Wait() // ожидаем завершения обоих горутин
	fmt.Println("Горутины завершили выполнение")
}

func solutionForGroup(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Горутина %d начала выполнение \n", id)
	time.Sleep(2 * time.Second)
	fmt.Printf("Горутина %d завершила выполнение \n", id)
}
