package state

const DefaultStateDir = "/var/lib/ovmh"

// State is a state
type State struct {
	// Disks is a map of disks
	Disks map[string]*Disk
	// Images is a map of images
	Images map[string]*Image
	// Networks is a map of networks
	Networks map[string]*Network
	// VMs is a map of VMs
	VMs map[string]*VM
}
