package main

import (
	"context"
	"log"
	"runtime"
	"sync"
	"time"
)

// 假设每个do需要150ms, 那么100个do需要15s, 假设我们4核心, 那么协程运行就应该是15s/4core=3.5s
// canbuy的情况不一定, 最快100ms最慢以最耗时的check点+前面已经跑过的协程时间
func main1() {
	var start = time.Now()

	//check()
	canBuy()

	end := time.Since(start).String()
	log.Print(end)

	time.Sleep(time.Minute)
}

func check() {
	runtime.GOMAXPROCS(4)

	var checkChan = make(chan int, 512)
	var wait sync.WaitGroup
	var res []int
	ctx, _ := context.WithCancel(context.Background())

	for i := 0; i < 100; i++ {
		j := i
		wait.Add(1)
		go do(ctx, &wait, checkChan, j)
	}

	wait.Wait()
	close(checkChan)

	for v := range checkChan {
		res = append(res, v)
	}

	log.Print(len(res))
	log.Print(res)
}

// 通道无数据
// 通道有数据
func canBuy() {
	runtime.GOMAXPROCS(4)

	var checkChan = make(chan int)
	defer close(checkChan)
	var wait sync.WaitGroup
	var res = int(999)
	ctx, cancel := context.WithCancel(context.Background())

	for i := 0; i < 100; i++ {
		j := i
		wait.Add(1)
		go do2(ctx, &wait, checkChan, j)
	}

	go func() {
		wait.Wait()
		cancel()
	}()

	go func() {
		for v := range checkChan {
			res = v
			cancel()
			return
		}
	}()

	select {
	case <-ctx.Done():
		log.Print(res)
	}
}

func do(ctx context.Context, wait *sync.WaitGroup, checkChan chan int, j int) {
	defer func() {
		if err := recover(); err != nil {
			log.Print(err)
		}
	}()
	defer wait.Done()
	select {
	case <-ctx.Done():
		return
	default:
		var d int
		for i := 0; i < 100000000; i++ {
			d += i
		}
		if j > 100 {
			var data = j
			checkChan <- data
		}
		return
	}
}

func do2(ctx context.Context, wait *sync.WaitGroup, checkChan chan int, j int) {
	defer func() {
		if err := recover(); err != nil {
			log.Print(err)
		}
	}()
	defer wait.Done()
	select {
	case <-ctx.Done():
		return
	default:
		var d int
		for i := 0; i < 100000000; i++ {
			d += i
		}
		log.Print(j)
		if j >100 {
			var data = j
			checkChan <- data
		}
		return
	}
}
