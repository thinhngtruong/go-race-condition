package main

import (
	"fmt"
	"sync"
)

var x int64 = 0
var mutex = &sync.Mutex{}

func addOne(wg *sync.WaitGroup) {
	mutex.Lock()
	x++
	mutex.Unlock()

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
