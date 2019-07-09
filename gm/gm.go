package gm

import (
	"fmt"
	"log"
	"strings"

	"github.com/go-mach/gm/config"
)

// Machinery is the main framework structure.
type Machinery struct {
	// gears []Gear
	gears map[string]Gear
}

// New initialize and return the main Machinery instance.
func New() *Machinery {
	return &Machinery{gears: make(map[string]Gear)}
}

// RegisterGear and configure a Gear with the Machinery.
func (m *Machinery) RegisterGear(gear Gear) *Machinery {
	if m.gears[gear.Name()] != nil {
		log.Fatalf("Gear %s already registered", gear.Name())
	} else {
		m.gears[gear.Name()] = gear
	}

	return m
}

// Register and configure multiple Gears with the Machinery.
func (m *Machinery) Register(gears ...Gear) *Machinery {
	for _, gear := range gears {
		m.RegisterGear(gear)
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
		go gear.Start(m)
	}
}

// GetGear returns a Gear instance pointer
// TODO: use a map to store Gears
func (m *Machinery) GetGear(name string) Gear {
	return m.gears[name]
}
