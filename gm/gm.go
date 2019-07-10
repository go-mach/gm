package gm

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/go-mach/gm/config"
)

// Machinery is the main framework structure.
type Machinery struct {
	gears        map[string]Gear
	GracefulStop chan os.Signal
}

// NewMachinery initialize and return the main Machinery engine instance.
func NewMachinery() *Machinery {
	theGoMachinery := &Machinery{
		gears:        make(map[string]Gear),
		GracefulStop: make(chan os.Signal),
	}

	signal.Notify(theGoMachinery.GracefulStop, syscall.SIGTERM)
	signal.Notify(theGoMachinery.GracefulStop, syscall.SIGINT)

	go func() {
		sig := <-theGoMachinery.GracefulStop
		log.Printf("caught sig: %+v", sig)
		log.Println("Wait for 2 second to finish processing")
		time.Sleep(2 * time.Second)
		os.Exit(0)
	}()

	return theGoMachinery
}

// With and configure one or more Gears with the Machinery engine.
func (m *Machinery) With(gears ...Gear) *Machinery {
	var gearName string

	for _, gear := range gears {
		gearName = gear.Name()
		if m.gears[gearName] != nil {
			log.Printf("Gear %s already registered", gearName)
		} else {
			log.Printf("registering %s Gear", gearName)
			m.gears[gearName] = gear
		}
	}

	return m
}

// Start configure app gears and starts the machinery
func (m *Machinery) Start() {
	log.Println("configuring machinery gears")
	m.configureGears()

	log.Println("starting machinery gears")
	m.startGears()

	log.Println("app Machinery started")
	select {}
}

// configure configurable gears
func (m *Machinery) configureGears() {
	for gearName, gear := range m.gears {
		// check if the gear is Configurable
		if configurableGear, ok := gear.(Configurable); ok {
			log.Printf("the %s gear is configurable", gearName)
			gearConfig := config.Get(strings.ToLower(gearName))
			if gearConfig == nil {
				panic(fmt.Sprintf("no configuration found for gear %s", gearName))
			}
			configurableGear.Configure(config.Get(gearName))
		}
	}
}

func (m *Machinery) startGears() {
	for gearName, gear := range m.gears {
		log.Printf("starting the %s gear", gearName)
		gear.Start(m)
	}
}

// GetGear returns a Gear instance pointer
// TODO: use a map to store Gears
func (m *Machinery) GetGear(name string) Gear {
	return m.gears[name]
}
