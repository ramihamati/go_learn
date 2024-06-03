package main

import (
	"log"
	"time"
)

func Channels1Start() {

	c := make(chan string)
	defer close(c)

	log.Println("Hello")

	go func() {
		for {
			message := <-c
			log.Println(message)
		}

	}()

	go func() {
		time.Sleep(2 * time.Second)
		c <- "what"
		log.Println("sent")
		c <- "what"
		c <- "what"
	}()

	time.Sleep(3 * time.Second)
}
