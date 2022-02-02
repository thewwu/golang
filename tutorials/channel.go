package main

import (
	"fmt"
	"sync"
)

func main() {
	channel_hardcode()
	channel_iterate()
}

////////////////////////////////////////////////////////////////////////
// use channel to send and receive values from concurrent go routine
////////////////////////////////////////////////////////////////////////
func foo(c chan int, someValue int) {
	c <- someValue * 5
}

func channel_hardcode() {
	fooVal := make(chan int)
	go foo(fooVal, 3)
	go foo(fooVal, 5)
	// v1 := <-fooVal
	// v2 := <-fooVal

	v1, v2 := <-fooVal, <-fooVal
	fmt.Println(v1, v2)
}

////////////////////////////////////////////////////////////////////////
// iterate the channel
////////////////////////////////////////////////////////////////////////
var wg sync.WaitGroup

func foo_sync(c chan int, someValue int) {
	defer wg.Done()
	c <- someValue * 5
}

func channel_iterate() {
	fmt.Println("\nIterate channel items...")
	fooVal := make(chan int, 10)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go foo_sync(fooVal, i)
	}
	wg.Wait()
	close(fooVal)

	for item := range fooVal {
		fmt.Println(item)
	}
}
