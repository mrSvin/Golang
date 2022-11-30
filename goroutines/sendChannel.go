package main

import (
	"fmt"
	"time"
)

func SendChannel() {
	intCh := make(chan int)
	go sender(intCh)
	go receiver(intCh)
	time.Sleep(5 * time.Second)

}

func sender(chanIn chan int) {
	defer close(chanIn)
	for i := 0; i < 10; i++ {
		fmt.Println("sender: ", i)
		chanIn <- i
	}
}

func receiver(chanIn chan int) {
	for {
		num, opened := <-chanIn // получаем данные из потока
		if !opened {
			break // если поток закрыт, выход из цикла
		}
		fmt.Println("receiver: ", num)
	}

}
