package main

import (
	"fmt"
	"rsc.io/quote"
)

func hello() {
	fmt.Println("Hello!")
	message := quote.Hello()
	fmt.Println(message)
}
