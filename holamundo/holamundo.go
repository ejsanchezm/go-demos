package main

import (
	//"rsc.io/quote"
	"fmt"
	"greetings"
)

func main() {
	//Code needed for calling the quote module
	//fmt.Println("Hola Mundo!")
	//fmt.Println(quote.Go())

	//Code needed for calling the greetings module
	message:= greetings.Hello("Eduardo")
	fmt.Println(message)
}

