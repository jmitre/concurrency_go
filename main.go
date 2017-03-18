package main

import (
	"github.com/jmitre/concurrency_go/Formatter"
	"time"
)

var threadPrinter *formatter.Formatter

func main() {
	threadPrinter = formatter.CreateFormatter()

	hasPrintedHelloChan := make(chan bool)
	hasPrintedWorldChan := make(chan bool)

	threadPrinter.PrintMainThread("go printHello()")
	go printHello(hasPrintedHelloChan, hasPrintedWorldChan) // starting thread to run printHello
	threadPrinter.PrintMainThread("go printWorld()")
	go printWorld(hasPrintedHelloChan, hasPrintedWorldChan) // starting thread to run printWolrd

	threadPrinter.PrintMainThread("hasPrintedWorldChan <- true")
	hasPrintedWorldChan <- true

	threadPrinter.PrintMainThread("time.Sleep(100 * time.Second)")
	time.Sleep(100 * time.Second)
}

func printHello(hasPrintedHelloChan chan bool, hasPrintedWolrdChan chan bool) {
	threadPrinter.PrintThreadOne("for i := 0; i < 10; i++ {")
	for i := 0; i < 10; i++ {
		threadPrinter.PrintThreadOne("<-hasPrintedWolrdChan")
		<-hasPrintedWolrdChan
		threadPrinter.PrintThreadOne("Hello")
		threadPrinter.PrintThreadOne("hasPrintedHelloChan <- true")
		hasPrintedHelloChan <- true
	}
}

func printWorld(hasPrintedHelloChan chan bool, hasPrintedWolrdChan chan bool) {
	threadPrinter.PrintThreadTwo("for i := 0; i < 10; i++ {")
	for i := 0; i < 10; i++ {
		threadPrinter.PrintThreadTwo("<-hasPrintedHelloChan")
		<-hasPrintedHelloChan
		threadPrinter.PrintThreadTwo("World")
		threadPrinter.PrintThreadTwo("hasPrintedWolrdChan <- true")
		hasPrintedWolrdChan <- true
	}
}
