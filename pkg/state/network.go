package state

// Network is a network
type Network struct {
	// Name is the name of the network
	Name string
	// Bridge is the bridge name
	Bridge string
	// Subnet is the subnet
	Subnet string
	// Gateway is the gateway
	Gateway string
	// DNS is the DNS server
	DNS string
	// Domain is the domain
	Domain string
	// VLAN is the VLAN
	VLAN int
}
