package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var x int64 = 0

func addOne(wg *sync.WaitGroup) {
	atomic.AddInt64(&x, 1)
	wg.Done()
}

func main() {
	var w sync.WaitGroup
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go addOne(&w)
	}
	w.Wait()
	fmt.Println(x)
}
