package formatter

import (
	"fmt"
	"os"
	"os/exec"
)

var SIZE = 20

type Formatter struct {
	PanelOne []string
	PanelTwo []string
}

func CreateFormatter() *Formatter {
	return &Formatter{make([]string, SIZE), make([]string, SIZE)}
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

func (f *Formatter) Print() {
	clear()
	for i := 0; i < SIZE; i++ {
		fmt.Printf("%-15s | %-15s\n", f.PanelOne[i], f.PanelTwo[i])
	}
	fmt.Println(f)
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
