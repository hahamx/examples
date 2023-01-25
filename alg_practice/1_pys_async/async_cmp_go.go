package main

import (
	"fmt"
	"os"
	"sync"
)

var (
	AsyncQueue = make(chan int, 2)
	wg         = sync.WaitGroup{}
)

func SyncPutin(n int) {
	AsyncQueue = make(chan int, n)
	for i := 0; i < n; i++ {
		AsyncQueue <- i
	}
}

func PutsIn(n int) {
	go func() {
		for i := 0; i < n; i++ {
			AsyncQueue <- i
		}

	}()

}

func Customer(n int) {
	wg.Add(1)
	go func() {
		var count int
		for {
			fmt.Println("get async item")
			if count >= n {
				wg.Done()
				os.Exit(1)
			}
			count += 1
			newOne := <-AsyncQueue
			fmt.Println("Async Got:", newOne)
		}
	}()
	wg.Wait()
}

// 同步存，异步取
func SyncCmp(times int) {

	SyncPutin(times)
	Customer(times)
}

// 异步存取
func AsyncCmp(times int) {

	PutsIn(times)

	Customer(times)
}

func main() {

	times := 10
	SyncCmp(times)
	AsyncCmp(times)

}
