package main

import "fmt"

func NoBufferChan() {
	intCh := make(chan int)
	go factorialChannel(6, intCh)
	fmt.Println(<-intCh) // получение данных из канала
	fmt.Println("The End")
}

func BufferChan() {

	intCh := make(chan int, 3)
	intCh <- 10
	intCh <- 2
	intCh <- 3
	//intCh <- 1 //блокировка - функция main ждет, когда освободится место в канале

	fmt.Println("Read after write 3 channel:")
	fmt.Println("cap:", cap(intCh))
	fmt.Println("length:", len(intCh))

	//После прочтения из буфера удаляются в порядке очереди
	fmt.Println("Read after read 2 channel:")
	val := <-intCh
	val2 := <-intCh
	fmt.Println("Read values: ", val, ",", val2)
	fmt.Println("cap:", cap(intCh))
	fmt.Println("length:", len(intCh))

	fmt.Println("The End")

}

func factorialChannel(n int, ch chan int) {

	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	fmt.Println(n, "-", result)
	ch <- result // отправка данных в канал
}

func OnlyReadOnlyWriteChannels() {
	//var inCh chan<- int        //only write
	//var outCh <-chan int //only read

	intCh := make(chan int)
	go factorialInReadOutWrite(5, intCh)
	fmt.Println("OnlyReadOnlyWriteChannels")
	fmt.Println(<-intCh)
	fmt.Println("The End")

}

func factorialInReadOutWrite(n int, ch chan<- int) {

	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	ch <- result
}

func ReturnChannel() {
	fmt.Println("ReturnChannel")
	if <-createChan(3) == 3 {
		fmt.Println("ok")
	}

}

func createChan(n int) chan int {
	ch := make(chan int) // создаем канал
	go func() {
		ch <- n // отправляем данные в канал
	}() // запускаем горутину
	return ch // возвращаем канал
}

func CloseChannel() {
	intCh := make(chan int, 3)
	intCh <- 10
	intCh <- 3
	intCh <- 1

	val, opened := <-intCh
	fmt.Println(val)
	fmt.Println(opened)

	val2, opened := <-intCh
	fmt.Println(val2)
	fmt.Println(opened)

	val3, opened := <-intCh
	fmt.Println(val3)
	fmt.Println(opened)

	close(intCh)

	val4, opened := <-intCh
	fmt.Println(val4)
	fmt.Println(opened)

	val5, opened := <-intCh
	fmt.Println(val5)
	fmt.Println(opened)

	val6, opened := <-intCh
	fmt.Println(val6)
	fmt.Println(opened)

	if opened == false {
		fmt.Println("Channel closed")
	}

}
