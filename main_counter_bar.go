package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func MainCounterBar() {
	var d = make(chan int64)
	go ProgressBar(d, 2000)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			d <- 10
			time.Sleep(time.Millisecond * 200)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			d <- 10
			time.Sleep(time.Millisecond * 400)
		}
	}()
	wg.Wait()
}

func bar(current, count int64, l ...int64) string {
	var size int64
	if len(l) > 0 {
		if l[0] == 0 {
			size = 100
		} else {
			size = l[0]
		}
	} else {
		size = 100
	}

	if current == 0 {
		str := ""
		for i := int64(0); i < size; i++ {
			str += " "
		}
		return "[" + str + "] 0%"
	}

	if current >= count || count == 0 {
		str := ""
		for i := int64(0); i < size; i++ {
			str += "="
		}
		return "[" + str + "] 100%"
	}

	percent := int64((float64(current) / float64(count)) * 100)
	currentEqual := int64((float64(current) / float64(count)) * float64(size))
	str := ""
	for i := int64(0); i < size; i++ {
		if i < currentEqual {
			str += "="
		} else {
			str += " "
		}
	}
	return "[" + str + "] " + strconv.Itoa(int(percent)) + "%"
}

// ProgressBar 打印进度条
func ProgressBar(c chan int64, count int64) {
	var current int64
	for number := range c {
		current += number
		str := bar(current, count, 50)
		fmt.Printf("\r%s %d/%d", str, current, count)
	}
}
