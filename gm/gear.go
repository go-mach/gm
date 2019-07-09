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
// However, a gear will not be configured if it does not implement
// the Configurable interface.
type BaseGear struct {
	Config map[string]interface{}
}
