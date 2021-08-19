package main

import (
	"log"
	"time"
)

func MainChannel() {
	var a = make(chan int, 10)
	a <- 1
	close(a)
	b1, ok1 := <-a
	log.Println(b1, ok1)
	b2, ok2 := <-a
	log.Println(b2, ok2)

	time.Sleep(time.Second * 5)
}
