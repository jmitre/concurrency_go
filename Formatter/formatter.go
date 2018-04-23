package formatter

import (
	"fmt"
	"os"
	"os/exec"
	"sync"
	"time"
)

var SIZE = 45

type Formatter struct {
	PanelOne   []string
	PanelTwo   []string
	PanelThree []string
	Step       int
	Mutex      *sync.Mutex
}

func CreateFormatter() *Formatter {
	return &Formatter{make([]string, SIZE), make([]string, SIZE), make([]string, SIZE), 0, &sync.Mutex{}}
}

func (f *Formatter) PrintMainThread(s string) {
	f.Mutex.Lock()
	f.AddToPanelOne(s)
	f.AddToPanelTwo(" ")
	f.AddToPanelThree(" ")
	f.Print()
	f.Mutex.Unlock()
}

func (f *Formatter) PrintThreadOne(s string) {
	f.Mutex.Lock()
	f.AddToPanelTwo(s)
	f.AddToPanelOne(" ")
	f.AddToPanelThree(" ")
	f.Print()
	f.Mutex.Unlock()
}

func (f *Formatter) PrintThreadTwo(s string) {
	f.Mutex.Lock()
	f.AddToPanelThree(s)
	f.AddToPanelOne(" ")
	f.AddToPanelTwo(" ")
	f.Print()
	f.Mutex.Unlock()
}

func (f *Formatter) AddToPanelOne(s string) {
	for i := 0; i < SIZE; i++ {
		if f.PanelOne[i] == "" {
			f.PanelOne[i] = s
			return
		}
	}
}

func (f *Formatter) AddToPanelTwo(s string) {
	for i := 0; i < SIZE; i++ {
		if f.PanelTwo[i] == "" {
			f.PanelTwo[i] = s
			return
		}
	}
}

func (f *Formatter) AddToPanelThree(s string) {
	for i := 0; i < SIZE; i++ {
		if f.PanelThree[i] == "" {
			f.PanelThree[i] = s
			return
		}
	}
}

func (f *Formatter) Print() {
	time.Sleep(2 * time.Second)
	clear()
	f.Step++
	fmt.Println("Step", f.Step)
	fmt.Printf("%-40s|%-40s|%-40s\n", "Main", "Thread 1", "Thread 2")
	fmt.Printf("%-40s|%-40s|%-40s\n", delimiter, delimiter, delimiter)
	for i := 0; i < SIZE; i++ {
		fmt.Printf("%-40s|%-40s|%-40s\n", f.PanelOne[i], f.PanelTwo[i], f.PanelThree[i])
	}
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

var delimiter = "########################################"
