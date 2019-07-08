package gm

import "log"

// Machinery is the main framework structure.
type Machinery struct {
	gears []Gear
}

// New initialize and return the main Machinery instance.
func New() *Machinery {
	return &Machinery{}
}

// Register and configure a Gear with the Machinery.
func (m *Machinery) Register(gear Gear) {
	m.gears = append(m.gears, gear)
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
	for _, gear := range m.gears {
		// check if the gear is Configurable
		if configurableGear, ok := gear.(Configurable); ok {
			log.Printf("the %s gear is configurable", gear.Name())
			configurableGear.Configure(nil)
		}
	}
}

func (m *Machinery) startGears() {
	for _, gear := range m.gears {
		log.Printf("starting the %s gear", gear.Name())
		go gear.Start(m)
	}
}
