package main

import (
	"fmt"
	"log"
	"rsc.io/quote"
)

func main() {
	log.Println("Hello", quote.Hello())
	fmt.Println(quote.Hello())
}
