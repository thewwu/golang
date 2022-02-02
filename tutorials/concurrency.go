package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	simple_go_routine()
	sync_go_routine()
	recover_go_routine()
}

func say(s string) {
	for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

// concurrency -> dealing multiple thing at 1 time -> single thread but using the wait time to complete another task
// parallel -> doing multiple thing at 1 time -> multi-thread / processing
// changing to go routine by adding "go" before the func
func simple_go_routine() {
	go say("Hey")
	say("there")

	// if 2 go -> nothing as all compulsory routine has run -> program ends before go routine
	// need to add a compulsory routine which process time is longer than the go routine or use sync
	go say("Hey")
	go say("Again")

	time.Sleep(time.Second)
}

////////////////////////////////////////////////////////////////
// Sync Go routine
////////////////////////////////////////////////////////////////
var wg sync.WaitGroup

func say_sync(s string) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
	// wg.Done() // deadlock if without this Done or use defer
}

func sync_go_routine() {
	fmt.Println("\nSync go routine...")

	wg.Add(1)
	go say_sync("Start")
	wg.Add(1)
	go say_sync("Again")
	wg.Wait()
}

////////////////////////////////////////////////////////////////
// Panic -> raise Exception, Recover -> try except
////////////////////////////////////////////////////////////////
func cleanup() {
	defer wg.Done()
	if r := recover(); r != nil {
		fmt.Println("Recovered in cleanup:", r)
	}
}

func say_recover(s string) {
	defer cleanup()
	for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
		if i == 2 {
			panic("Oh dear... a 2")
		}
	}
}

func recover_go_routine() {
	fmt.Println("\nRecover go routine...")

	wg.Add(1)
	go say_recover("Start")
	wg.Add(1)
	go say_recover("Again")
	wg.Wait()
}
