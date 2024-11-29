package state

const DefaultStateDir = "/var/lib/ovmh"

// State is a state
type State struct {
	// Path is the path to the state directory
	Path string
	// VMs is a map of VMs
	VMs map[string]*VM
}
