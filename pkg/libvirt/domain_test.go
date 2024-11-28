package libvirt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const testDomainXML = `
<domain type="kvm">
  <name>archlinux</name>
  <uuid>328c4a91-565d-4e43-860c-a4cfcb007293</uuid>
  <metadata>
    <libosinfo:libosinfo xmlns:libosinfo="http://libosinfo.org/xmlns/libvirt/domain/1.0">
      <libosinfo:os id="http://archlinux.org/archlinux/rolling"/>
    </libosinfo:libosinfo>
  </metadata>
  <memory unit="KiB">16777216</memory>
  <currentMemory unit="KiB">16777216</currentMemory>
  <vcpu placement="static">16</vcpu>
  <os>
    <type arch="x86_64" machine="pc-q35-8.2">hvm</type>
    <boot dev="hd"/>
  </os>
  <features>
    <acpi/>
    <apic/>
  </features>
  <cpu mode="host-passthrough" check="none" migratable="on"/>
  <clock offset="utc">
    <timer name="rtc" tickpolicy="catchup"/>
    <timer name="pit" tickpolicy="delay"/>
    <timer name="hpet" present="no"/>
  </clock>
  <on_poweroff>destroy</on_poweroff>
  <on_reboot>restart</on_reboot>
  <on_crash>destroy</on_crash>
  <pm>
    <suspend-to-mem enabled="no"/>
    <suspend-to-disk enabled="no"/>
  </pm>
  <devices>
    <emulator>/usr/bin/qemu-system-x86_64</emulator>
    <disk type="block" device="disk">
      <driver name="qemu" type="raw" cache="none" io="native" discard="unmap"/>
      <source dev="/dev/ssdstore/archlinux"/>
      <target dev="vda" bus="virtio"/>
      <address type="pci" domain="0x0000" bus="0x04" slot="0x00" function="0x0"/>
    </disk>
    <disk type="file" device="cdrom">
      <driver name="qemu" type="raw"/>
      <target dev="sda" bus="sata"/>
      <readonly/>
      <address type="drive" controller="0" bus="0" target="0" unit="0"/>
    </disk>
    <controller type="usb" index="0" model="qemu-xhci" ports="15">
      <address type="pci" domain="0x0000" bus="0x02" slot="0x00" function="0x0"/>
    </controller>
    <controller type="pci" index="0" model="pcie-root"/>
    <controller type="pci" index="1" model="pcie-root-port">
      <model name="pcie-root-port"/>
      <target chassis="1" port="0x10"/>
      <address type="pci" domain="0x0000" bus="0x00" slot="0x02" function="0x0" multifunction="on"/>
    </controller>
    <controller type="pci" index="2" model="pcie-root-port">
      <model name="pcie-root-port"/>
      <target chassis="2" port="0x11"/>
      <address type="pci" domain="0x0000" bus="0x00" slot="0x02" function="0x1"/>
    </controller>
    <controller type="pci" index="3" model="pcie-root-port">
      <model name="pcie-root-port"/>
      <target chassis="3" port="0x12"/>
      <address type="pci" domain="0x0000" bus="0x00" slot="0x02" function="0x2"/>
    </controller>
    <controller type="pci" index="4" model="pcie-root-port">
      <model name="pcie-root-port"/>
      <target chassis="4" port="0x13"/>
      <address type="pci" domain="0x0000" bus="0x00" slot="0x02" function="0x3"/>
    </controller>
    <controller type="pci" index="5" model="pcie-root-port">
      <model name="pcie-root-port"/>
      <target chassis="5" port="0x14"/>
      <address type="pci" domain="0x0000" bus="0x00" slot="0x02" function="0x4"/>
    </controller>
    <controller type="pci" index="6" model="pcie-root-port">
      <model name="pcie-root-port"/>
      <target chassis="6" port="0x15"/>
      <address type="pci" domain="0x0000" bus="0x00" slot="0x02" function="0x5"/>
    </controller>
    <controller type="pci" index="7" model="pcie-root-port">
      <model name="pcie-root-port"/>
      <target chassis="7" port="0x16"/>
      <address type="pci" domain="0x0000" bus="0x00" slot="0x02" function="0x6"/>
    </controller>
    <controller type="pci" index="8" model="pcie-root-port">
      <model name="pcie-root-port"/>
      <target chassis="8" port="0x17"/>
      <address type="pci" domain="0x0000" bus="0x00" slot="0x02" function="0x7"/>
    </controller>
    <controller type="pci" index="9" model="pcie-root-port">
      <model name="pcie-root-port"/>
      <target chassis="9" port="0x18"/>
      <address type="pci" domain="0x0000" bus="0x00" slot="0x03" function="0x0" multifunction="on"/>
    </controller>
    <controller type="pci" index="10" model="pcie-root-port">
      <model name="pcie-root-port"/>
      <target chassis="10" port="0x19"/>
      <address type="pci" domain="0x0000" bus="0x00" slot="0x03" function="0x1"/>
    </controller>
    <controller type="pci" index="11" model="pcie-root-port">
      <model name="pcie-root-port"/>
      <target chassis="11" port="0x1a"/>
      <address type="pci" domain="0x0000" bus="0x00" slot="0x03" function="0x2"/>
    </controller>
    <controller type="pci" index="12" model="pcie-root-port">
      <model name="pcie-root-port"/>
      <target chassis="12" port="0x1b"/>
      <address type="pci" domain="0x0000" bus="0x00" slot="0x03" function="0x3"/>
    </controller>
    <controller type="pci" index="13" model="pcie-root-port">
      <model name="pcie-root-port"/>
      <target chassis="13" port="0x1c"/>
      <address type="pci" domain="0x0000" bus="0x00" slot="0x03" function="0x4"/>
    </controller>
    <controller type="pci" index="14" model="pcie-root-port">
      <model name="pcie-root-port"/>
      <target chassis="14" port="0x1d"/>
      <address type="pci" domain="0x0000" bus="0x00" slot="0x03" function="0x5"/>
    </controller>
    <controller type="sata" index="0">
      <address type="pci" domain="0x0000" bus="0x00" slot="0x1f" function="0x2"/>
    </controller>
    <controller type="virtio-serial" index="0">
      <address type="pci" domain="0x0000" bus="0x03" slot="0x00" function="0x0"/>
    </controller>
    <interface type="bridge">
      <mac address="52:54:00:af:f7:af"/>
      <source bridge="br0"/>
      <vlan>
        <tag id="8"/>
      </vlan>
      <virtualport type="openvswitch">
        <parameters interfaceid="f3540199-4cb6-4f33-81f9-6c6a381c6993"/>
      </virtualport>
      <model type="virtio"/>
      <address type="pci" domain="0x0000" bus="0x01" slot="0x00" function="0x0"/>
    </interface>
    <serial type="pty">
      <target type="isa-serial" port="0">
        <model name="isa-serial"/>
      </target>
    </serial>
    <console type="pty">
      <target type="serial" port="0"/>
    </console>
    <channel type="unix">
      <target type="virtio" name="org.qemu.guest_agent.0"/>
      <address type="virtio-serial" controller="0" bus="0" port="1"/>
    </channel>
    <input type="tablet" bus="usb">
      <address type="usb" bus="0" port="1"/>
    </input>
    <input type="mouse" bus="ps2"/>
    <input type="keyboard" bus="ps2"/>
    <graphics type="vnc" port="-1" autoport="yes">
      <listen type="address"/>
    </graphics>
    <audio id="1" type="none"/>
    <video>
      <model type="virtio" heads="1" primary="yes"/>
      <address type="pci" domain="0x0000" bus="0x00" slot="0x01" function="0x0"/>
    </video>
    <watchdog model="itco" action="reset"/>
    <memballoon model="virtio">
      <address type="pci" domain="0x0000" bus="0x05" slot="0x00" function="0x0"/>
    </memballoon>
    <rng model="virtio">
      <backend model="random">/dev/urandom</backend>
      <address type="pci" domain="0x0000" bus="0x06" slot="0x00" function="0x0"/>
    </rng>
  </devices>
</domain>
`

const testDomainXML2 = `<domain type="kvm">
  <name>archlinux</name>
  <uuid>328c4a91-565d-4e43-860c-a4cfcb007293</uuid>
  <metadata>
    <libosinfo xmlns="http://libosinfo.org/xmlns/libvirt/domain/1.0">
      <os xmlns="http://libosinfo.org/xmlns/libvirt/domain/1.0" id="http://archlinux.org/archlinux/rolling"></os>
    </libosinfo>
  </metadata>
  <memory unit="KiB">16777216</memory>
  <vcpu placement="static">16</vcpu>
  <os>
    <type arch="x86_64" machine="pc-q35-8.2">hvm</type>
    <boot dev="hd"></boot>
  </os>
  <features>
    <acpi></acpi>
    <apic></apic>
  </features>
  <cpu mode="host-passthrough" check="none" migratable="on"></cpu>
  <clock offset="utc">
    <timer name="rtc" tickpolicy="catchup"></timer>
    <timer name="pit" tickpolicy="delay"></timer>
    <timer name="hpet" present="no"></timer>
  </clock>
  <on_poweroff>destroy</on_poweroff>
  <on_reboot>restart</on_reboot>
  <on_crash>destroy</on_crash>
  <pm>
    <suspend-to-mem enabled="no"></suspend-to-mem>
    <suspend-to-disk enabled="no"></suspend-to-disk>
  </pm>
  <devices>
    <emulator>/usr/bin/qemu-system-x86_64</emulator>
    <disk type="block" device="disk">
      <driver name="qemu" type="raw" cache="none" io="native" discard="unmap"></driver>
      <source dev="/dev/ssdstore/archlinux"></source>
      <target dev="vda" bus="virtio"></target>
    </disk>
    <interface type="bridge">
      <mac address="52:54:00:af:f7:af"></mac>
      <source bridge="br0"></source>
      <vlan>
        <tag id="8"></tag>
      </vlan>
      <virtualport type="openvswitch"></virtualport>
      <model name="virtio"></model>
    </interface>
    <serial type="pty">
      <target type="isa-serial" port="0">
        <model name="isa-serial"></model>
      </target>
    </serial>
    <console type="pty">
      <target type="serial" port="0"></target>
    </console>
    <channel type="unix">
      <target type="virtio" name="org.qemu.guest_agent.0"></target>
    </channel>
    <graphics type="vnc" port="-1" autoport="yes">
      <listen type="address"></listen>
    </graphics>
    <audio id="1" type="none"></audio>
    <video>
      <model type="virtio" heads="1" primary="yes"></model>
    </video>
    <watchdog model="itco" action="reset"></watchdog>
    <rng model="virtio">
      <backend model="random">/dev/urandom</backend>
    </rng>
  </devices>
</domain>`

func TestUnmarshalDomain(t *testing.T) {
	domain, err := UnmarshalDomain([]byte(testDomainXML))
	assert.Nil(t, err)
	assert.Equal(t, "archlinux", domain.Name)
	assert.Equal(t, "328c4a91-565d-4e43-860c-a4cfcb007293", domain.UUID)
}

func TestMarshalDomain(t *testing.T) {
	domain := &Domain{
		Type: "kvm",
		Name: "archlinux",
		UUID: "328c4a91-565d-4e43-860c-a4cfcb007293",
		Metadata: Metadata{
			Libosinfo: Libosinfo{
				OS: OsInfoOS{
					ID: "http://archlinux.org/archlinux/rolling",
				},
			},
		},
		OS: OS{
			Type: OSType{
				Arch:    "x86_64",
				Machine: "pc-q35-8.2",
				Type:    "hvm",
			},
			Boot: Boot{
				Dev: "hd",
			},
		},
		Memory: Memory{
			Unit:  "KiB",
			Value: 16777216,
		},
		VCPU: VCPU{
			Placement: "static",
			Value:     16,
		},
		Features: Features{
			ACPI: struct{}{},
			APIC: struct{}{},
		},
		CPU: CPU{
			Mode:       "host-passthrough",
			Check:      "none",
			Migratable: "on",
		},
		Clock: Clock{
			Offset: "utc",
			Timers: []Timer{
				{
					Name:       "rtc",
					TickPolicy: "catchup",
				},
				{
					Name:       "pit",
					TickPolicy: "delay",
				},
				{
					Name:    "hpet",
					Present: "no",
				},
			},
		},
		OnPoweroff: "destroy",
		OnReboot:   "restart",
		OnCrash:    "destroy",
		PM: PM{
			SuspendToMem:  SuspendOption{Enabled: "no"},
			SuspendToDisk: SuspendOption{Enabled: "no"},
		},
		Devices: Devices{
			Emulator: "/usr/bin/qemu-system-x86_64",
			Disks: []Disk{
				{
					Type:   "block",
					Device: "disk",
					Driver: &Driver{
						Name:    "qemu",
						Type:    "raw",
						Cache:   "none",
						IO:      "native",
						Discard: "unmap",
					},
					Source: &Source{
						Dev: "/dev/ssdstore/archlinux",
					},
					Target: &Target{
						Dev: "vda",
						Bus: "virtio",
					},
				},
			},
			Interfaces: []Interface{
				{
					Type: "bridge",
					MAC:  MAC{Address: "52:54:00:af:f7:af"},
					Source: InterfaceSource{
						Bridge: "br0",
					},
					VLAN: VLAN{
						Tag: VLANTag{ID: 8},
					},
					VirtualPort: VirtualPort{
						Type: "openvswitch",
					},
					Model: ModelElement{
						Name: "virtio",
					},
				},
			},
			Serial: &Serial{
				Type: "pty",
				Target: SerialTarget{
					Type:  "isa-serial",
					Port:  0,
					Model: &ModelElement{Name: "isa-serial"},
				},
			},
			Console: &Console{
				Type: "pty",
				Target: ConsoleTarget{
					Type: "serial",
					Port: 0,
				},
			},
			Channel: &Channel{
				Type: "unix",
				Target: ChannelTarget{
					Type: "virtio",
					Name: "org.qemu.guest_agent.0",
				},
			},
			Graphics: &Graphics{
				Type:     "vnc",
				Port:     -1,
				AutoPort: "yes",
				Listen: &Listen{
					Type: "address",
				},
			},
			Audio: &Audio{
				ID:   1,
				Type: "none",
			},
			Video: &Video{
				Model: VideoModel{
					Type:    "virtio",
					Heads:   1,
					Primary: "yes",
				},
			},
			Watchdog: &Watchdog{
				Model:  "itco",
				Action: "reset",
			},
			RNG: &RNG{
				Model: "virtio",
				Backend: RNGBackend{
					Model: "random",
					Path:  "/dev/urandom",
				},
			},
		},
	}
	xml, err := MarshalDomain(domain)
	assert.Nil(t, err)
	assert.Equal(t, testDomainXML2, string(xml))
}
