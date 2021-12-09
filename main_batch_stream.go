// 流批处理

package main

import (
	"context"
	"fmt"
	"time"
)

var Chan1 = make(chan int)
var Chan2 = make(chan []int)
var batchSize = 10
var batchTime = time.Second * 2

func MainBatchStream() {
	go consumerChan2()
	go batchDataStream()
	go producerChan1()

	select {}
}

func batchDataStream() {
	batchData := make([]int, 0, batchSize)
	ctx, cancel := context.WithTimeout(context.Background(), batchTime)
	for {
		select {
		case data, ok := <-Chan1:
			if !ok {
				continue
			}
			batchData = append(batchData, data)
			if len(batchData) >= batchSize {
				Chan2 <- batchData
				batchData = make([]int, 0, batchSize)
				cancel()
			}
		case <-ctx.Done():
			ctx, cancel = context.WithTimeout(context.Background(), batchTime)
			if len(batchData) > 0 {
				Chan2 <- batchData
				batchData = make([]int, 0, batchSize)
			}
		}
	}
}

func consumerChan2() {
	for d := range Chan2 {
		fmt.Printf("%d\r\n", d)
	}
}

func producerChan1() {
	for i := 0; i < 100; i++ {
		time.Sleep(200 * time.Millisecond)
		Chan1 <- i
	}
}
