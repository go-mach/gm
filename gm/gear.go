package gm

// type GearConfig struct {
// 	Name string
// }

// Configurable is the interface to mark a gear as configurable
type Configurable interface {
	Configure(config map[string]interface{})
}

// Gear is the Machinery main building block interface.
// If a component want to be loaded into the app have to implemet this interface.
type Gear interface {
	Name() string
	Start(machinery *Machinery)
}
