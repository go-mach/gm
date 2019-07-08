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
}

// BaseGear is the MAchinery main building block structure.
// If a component want to be loaded into the app should derive from this,
// or define Config map to store configuration
type BaseGear struct {
	config map[string]interface{}
}
