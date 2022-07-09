package main

import (
	"fmt"
	"sync"
)

func test1(wg *sync.WaitGroup) {
	fmt.Println("test1")
	wg.Done()
}

func test2(wg *sync.WaitGroup) {
	fmt.Println("test2")
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	test1(&wg)
	test2(&wg)

	wg.Wait()
	fmt.Println("Done")
}
