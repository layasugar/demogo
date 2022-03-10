package main

import "log"

func main() {

	var a = make(chan int)
	go func() {
		a <- 1
	}()

	c := <-a

	log.Print(c)

}
