package gm

// Configurable is the interface to mark a gear as configurable
type Configurable interface {
	Configure(config interface{})
}

// Gear is the Machinery main building block interface.
// If a component want to be loaded into the app have to implemet this interface.
type Gear interface {
	Name() string
	Start(machinery *Machinery)
	Provide() interface{}
}

// BaseGear is the Machinery most basic building block structure.
// If a component want to be loaded into the app should derive from this.
type BaseGear struct {
	UniqueName string
}

// ConfigurableGear is a BasicGear with a config map structure.
// However, a gear will not be configured if it does not implement
// the Configurable interface.
type ConfigurableGear struct {
	BaseGear
	Config map[string]interface{}
}
