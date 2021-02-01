package main

import (
	"fmt"
	"time"
)

func main() {
	// Go Routine!! Just attach go in front of Function
	// go count("Kim")
	// Caution, Go Routine only Runs During main runs
	// go count("Go")
	// time.Sleep(time.Millisecond * 500)

	// Channel Example
	// chan, Channel is interface btw main and Go Routine
	c := make(chan bool)
	people := [2]string{"Kim", "Park "}
	for _, name := range people {
		go isSexy(name, c)
	}
	fmt.Println(<-c)
	fmt.Println(<-c)
	// deadlock!!! Only twice times called
	// fmt.Println(<-c)
}

func count(name string) {
	for i := 0; i < 10; i++ {
		fmt.Println(name, "is fucking ", i, "times")
		time.Sleep(time.Millisecond * 100)
	}
}

func isSexy(name string, c chan bool) {
	fmt.Println(name)
	time.Sleep(time.Second * 1)
	c <- true
}
