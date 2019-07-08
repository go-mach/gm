package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-mach/gm/gm"
)

// SimpleGear .
type SimpleGear struct {
}

// Name .
func (sg *SimpleGear) Name() string {
	return "simple-gear"
}

// Start .
func (sg *SimpleGear) Start(m *gm.Machinery) {
	for {
		fmt.Println("Test")
		time.Sleep(2 * time.Second)
	}
}

// Configure .
func (sg *SimpleGear) Configure(config interface{}) {
	log.Printf("%s configured", sg.Name())
}

func main() {
	gm := gm.New()
	gm.Register(&SimpleGear{})
	gm.Start()
}
