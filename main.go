package main

import (
	"github.com/jmitre/concurrency_example/Formatter"
	"time"
)

func main() {
	formatter := formatter.CreateFormatter()
	formatter.AddToPanelOne("Hey")
	formatter.AddToPanelTwo("there")
	formatter.Print()
	time.Sleep(5 * time.Second)
	formatter.AddToPanelOne("Hey")
	formatter.AddToPanelTwo("there")
	formatter.AddToPanelOne("Hey")
	formatter.Print()
	time.Sleep(5 * time.Second)
	formatter.AddToPanelTwo("there")
	formatter.AddToPanelOne("Hey")
	formatter.AddToPanelTwo("there")
	formatter.AddToPanelOne("Hey")
	formatter.Print()
}
