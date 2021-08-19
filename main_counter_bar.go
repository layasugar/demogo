package main

import (
	"fmt"
	"strconv"
	"time"
)

func MainCounterBar() {
	// 总重建行数
	var count int64 = 10
	// 计数器
	var counter = make(chan int64)
	// 当前已重建的数据行数
	var current int64

	go func() {
		for n := range counter {
			str := bar(n, count, 50)
			fmt.Printf("\r%s", str)
		}
	}()

	for i := 1; i <= 10; i++ {
		current = int64(i)
		counter <- current
		time.Sleep(time.Second)
	}
	close(counter)
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
