package libvirt

import (
	"encoding/xml"
)

// Domain represents the root element of the XML.
type Domain struct {
	XMLName       xml.Name `xml:"domain"`
	Type          string   `xml:"type,attr"`
	Name          string   `xml:"name"`
	UUID          string   `xml:"uuid"`
	Metadata      Metadata `xml:"metadata"`
	Memory        Memory   `xml:"memory"`
	CurrentMemory *Memory  `xml:"currentMemory,omitempty"`
	VCPU          VCPU     `xml:"vcpu"`
	OS            OS       `xml:"os"`
	Features      Features `xml:"features"`
	CPU           CPU      `xml:"cpu"`
	Clock         Clock    `xml:"clock"`
	OnPoweroff    string   `xml:"on_poweroff"`
	OnReboot      string   `xml:"on_reboot"`
	OnCrash       string   `xml:"on_crash"`
	PM            PM       `xml:"pm"`
	Devices       Devices  `xml:"devices"`
}

// Metadata represents the <metadata> element.
type Metadata struct {
	Libosinfo Libosinfo `xml:"http://libosinfo.org/xmlns/libvirt/domain/1.0 libosinfo"`
}

// Libosinfo represents the <libosinfo:libosinfo> element.
type Libosinfo struct {
	OS OsInfoOS `xml:"http://libosinfo.org/xmlns/libvirt/domain/1.0 os"`
}

// OsInfoOS represents the <libosinfo:os> element.
type OsInfoOS struct {
	ID string `xml:"id,attr"`
}

// Memory represents the <memory> and <currentMemory> elements.
type Memory struct {
	Unit  string `xml:"unit,attr"`
	Value int    `xml:",chardata"`
}

// VCPU represents the <vcpu> element.
type VCPU struct {
	Placement string `xml:"placement,attr"`
	Value     int    `xml:",chardata"`
}

// OS represents the <os> element.
type OS struct {
	Type OSType `xml:"type"`
	Boot Boot   `xml:"boot"`
}

// OSType represents the <type> element within <os>.
type OSType struct {
	Arch    string `xml:"arch,attr"`
	Machine string `xml:"machine,attr"`
	Type    string `xml:",chardata"`
}

// Boot represents the <boot> element within <os>.
type Boot struct {
	Dev string `xml:"dev,attr"`
}

// Features represents the <features> element.
type Features struct {
	ACPI interface{} `xml:"acpi"`
	APIC interface{} `xml:"apic"`
}

// CPU represents the <cpu> element.
type CPU struct {
	Mode       string `xml:"mode,attr"`
	Check      string `xml:"check,attr"`
	Migratable string `xml:"migratable,attr"`
}

// Clock represents the <clock> element.
type Clock struct {
	Offset string  `xml:"offset,attr"`
	Timers []Timer `xml:"timer"`
}

// Timer represents the <timer> elements within <clock>.
type Timer struct {
	Name       string `xml:"name,attr"`
	TickPolicy string `xml:"tickpolicy,attr,omitempty"`
	Present    string `xml:"present,attr,omitempty"`
}

// PM represents the <pm> element.
type PM struct {
	SuspendToMem  SuspendOption `xml:"suspend-to-mem"`
	SuspendToDisk SuspendOption `xml:"suspend-to-disk"`
}

// SuspendOption represents the <suspend-to-mem> and <suspend-to-disk> elements.
type SuspendOption struct {
	Enabled string `xml:"enabled,attr"`
}

// Devices represents the <devices> element.
type Devices struct {
	Emulator    string       `xml:"emulator"`
	Disks       []Disk       `xml:"disk"`
	Controllers []Controller `xml:"controller"`
	Interfaces  []Interface  `xml:"interface"`
	Serial      *Serial      `xml:"serial,omitempty"`
	Console     *Console     `xml:"console,omitempty"`
	Channel     *Channel     `xml:"channel,omitempty"`
	Inputs      []Input      `xml:"input"`
	Graphics    *Graphics    `xml:"graphics,omitempty"`
	Audio       *Audio       `xml:"audio,omitempty"`
	Video       *Video       `xml:"video,omitempty"`
	Watchdog    *Watchdog    `xml:"watchdog,omitempty"`
	MemBalloon  *MemBalloon  `xml:"memballoon,omitempty"`
	RNG         *RNG         `xml:"rng,omitempty"`
}

// Disk represents the <disk> elements within <devices>.
type Disk struct {
	Type     string    `xml:"type,attr"`
	Device   string    `xml:"device,attr"`
	Driver   *Driver   `xml:"driver"`
	Source   *Source   `xml:"source"`
	Target   *Target   `xml:"target"`
	Address  *Address  `xml:"address"`
	ReadOnly *struct{} `xml:"readonly,omitempty"`
}

// Driver represents the <driver> element within <disk>.
type Driver struct {
	Name    string `xml:"name,attr,omitempty"`
	Type    string `xml:"type,attr,omitempty"`
	Cache   string `xml:"cache,attr,omitempty"`
	IO      string `xml:"io,attr,omitempty"`
	Discard string `xml:"discard,attr,omitempty"`
}

// Source represents the <source> element within <disk>.
type Source struct {
	Dev  string `xml:"dev,attr,omitempty"`
	File string `xml:"file,attr,omitempty"`
}

// Target represents the <target> element within <disk>.
type Target struct {
	Dev string `xml:"dev,attr"`
	Bus string `xml:"bus,attr"`
}

// Address represents the <address> elements.
type Address struct {
	Type          string `xml:"type,attr"`
	Domain        string `xml:"domain,attr,omitempty"`
	Bus           string `xml:"bus,attr,omitempty"`
	Slot          string `xml:"slot,attr,omitempty"`
	Function      string `xml:"function,attr,omitempty"`
	Controller    string `xml:"controller,attr,omitempty"`
	Target        string `xml:"target,attr,omitempty"`
	Unit          string `xml:"unit,attr,omitempty"`
	Port          string `xml:"port,attr,omitempty"`
	Multifunction string `xml:"multifunction,attr,omitempty"`
}

// Controller represents the <controller> elements within <devices>.
type Controller struct {
	Type         string         `xml:"type,attr"`
	Index        int            `xml:"index,attr"`
	Model        string         `xml:"model,attr"`
	Ports        int            `xml:"ports,attr,omitempty"`
	Address      *Address       `xml:"address"`
	ModelElement *ModelElement  `xml:"model"`
	Target       *TargetElement `xml:"target"`
}

// ModelElement represents the <model> element within <controller>.
type ModelElement struct {
	Name string `xml:"name,attr"`
}

// TargetElement represents the <target> element within <controller>.
type TargetElement struct {
	Chassis int    `xml:"chassis,attr"`
	Port    string `xml:"port,attr"`
}

// Interface represents the <interface> element within <devices>.
type Interface struct {
	Type        string          `xml:"type,attr"`
	MAC         MAC             `xml:"mac"`
	Source      InterfaceSource `xml:"source"`
	VLAN        VLAN            `xml:"vlan"`
	VirtualPort VirtualPort     `xml:"virtualport"`
	Model       ModelElement    `xml:"model"`
	Address     *Address        `xml:"address,omitempty"`
}

// MAC represents the <mac> element within <interface>.
type MAC struct {
	Address string `xml:"address,attr"`
}

// InterfaceSource represents the <source> element within <interface>.
type InterfaceSource struct {
	Bridge string `xml:"bridge,attr"`
}

// VLAN represents the <vlan> element within <interface>.
type VLAN struct {
	Tag VLANTag `xml:"tag"`
}

// VLANTag represents the <tag> element within <vlan>.
type VLANTag struct {
	ID int `xml:"id,attr"`
}

// VirtualPort represents the <virtualport> element within <interface>.
type VirtualPort struct {
	Type       string      `xml:"type,attr"`
	Parameters *Parameters `xml:"parameters"`
}

// Parameters represents the <parameters> element within <virtualport>.
type Parameters struct {
	InterfaceID string `xml:"interfaceid,attr"`
}

// Serial represents the <serial> element within <devices>.
type Serial struct {
	Type   string       `xml:"type,attr"`
	Target SerialTarget `xml:"target"`
}

// SerialTarget represents the <target> element within <serial>.
type SerialTarget struct {
	Type  string        `xml:"type,attr"`
	Port  int           `xml:"port,attr"`
	Model *ModelElement `xml:"model,omitempty"`
}

// Console represents the <console> element within <devices>.
type Console struct {
	Type   string        `xml:"type,attr"`
	Target ConsoleTarget `xml:"target"`
}

// ConsoleTarget represents the <target> element within <console>.
type ConsoleTarget struct {
	Type string `xml:"type,attr"`
	Port int    `xml:"port,attr"`
}

// Channel represents the <channel> element within <devices>.
type Channel struct {
	Type    string        `xml:"type,attr"`
	Target  ChannelTarget `xml:"target"`
	Address *Address      `xml:"address,omitempty"`
}

// ChannelTarget represents the <target> element within <channel>.
type ChannelTarget struct {
	Type string `xml:"type,attr"`
	Name string `xml:"name,attr"`
}

// Input represents the <input> elements within <devices>.
type Input struct {
	Type    string   `xml:"type,attr"`
	Bus     string   `xml:"bus,attr"`
	Address *Address `xml:"address,omitempty"`
}

// Graphics represents the <graphics> element within <devices>.
type Graphics struct {
	Type     string  `xml:"type,attr"`
	Port     int     `xml:"port,attr"`
	AutoPort string  `xml:"autoport,attr"`
	Listen   *Listen `xml:"listen"`
}

// Listen represents the <listen> element within <graphics>.
type Listen struct {
	Type string `xml:"type,attr"`
}

// Audio represents the <audio> element within <devices>.
type Audio struct {
	ID   int    `xml:"id,attr"`
	Type string `xml:"type,attr"`
}

// Video represents the <video> element within <devices>.
type Video struct {
	Model   VideoModel `xml:"model"`
	Address *Address   `xml:"address,omitempty"`
}

// VideoModel represents the <model> element within <video>.
type VideoModel struct {
	Type    string `xml:"type,attr"`
	Heads   int    `xml:"heads,attr"`
	Primary string `xml:"primary,attr"`
}

// Watchdog represents the <watchdog> element within <devices>.
type Watchdog struct {
	Model  string `xml:"model,attr"`
	Action string `xml:"action,attr"`
}

// MemBalloon represents the <memballoon> element within <devices>.
type MemBalloon struct {
	Model   string   `xml:"model,attr"`
	Address *Address `xml:"address"`
}

// RNG represents the <rng> element within <devices>.
type RNG struct {
	Model   string     `xml:"model,attr"`
	Backend RNGBackend `xml:"backend"`
	Address *Address   `xml:"address,omitempty"`
}

// RNGBackend represents the <backend> element within <rng>.
type RNGBackend struct {
	Model string `xml:"model,attr"`
	Path  string `xml:",chardata"`
}

// UnmarshalDomain unmarshals XML data into a Domain struct.
func UnmarshalDomain(data []byte) (*Domain, error) {
	var domain Domain
	err := xml.Unmarshal(data, &domain)
	if err != nil {
		return nil, err
	}
	return &domain, nil
}

// MarshalDomain marshals a Domain struct into XML data.
func MarshalDomain(domain *Domain) ([]byte, error) {
	return xml.MarshalIndent(domain, "", "  ")
}
