package main

import (
	"fmt"
	"time"
)

func main() {
	hasPrintedHelloChan := make(chan bool)
	hasPrintedWorldChan := make(chan bool)

	go printHello(hasPrintedHelloChan, hasPrintedWorldChan) // starting thread to run printHello
	go printWorld(hasPrintedHelloChan, hasPrintedWorldChan) // starting thread to run printWolrd

	hasPrintedWorldChan <- true

	time.Sleep(100 * time.Second)
}

func printHello(hasPrintedHelloChan chan bool, hasPrintedWolrdChan chan bool) {
	for i := 0; i < 10; i++ {
		<-hasPrintedWolrdChan
		fmt.Print("Hello ")
		hasPrintedHelloChan <- true
	}
}

func printWorld(hasPrintedHelloChan chan bool, hasPrintedWolrdChan chan bool) {
	for i := 0; i < 10; i++ {
		<-hasPrintedHelloChan
		fmt.Print("World\n")
		hasPrintedWolrdChan <- true
	}
}
