package libvirt

import (
	"libvirt.org/libvirt-go"
)

// Client is a libvirt client
type Client struct {
	uri  string
	conn *libvirt.Connect
}

// NewClient creates a new libvirt client
func NewClient(uri string) (*Client, error) {
	conn, err := libvirt.NewConnect(uri)
	if err != nil {
		return nil, err
	}

	return &Client{
		uri:  uri,
		conn: conn,
	}, nil
}

// Close closes the libvirt client
func (c *Client) Close() error {
	_, err := c.conn.Close()
	return err
}

// CreateDomain creates a new domain
func (c *Client) CreateDomain(domain *Domain) error {
	domainXML, err := MarshalDomain(domain)
	if err != nil {
		return err
	}

	_, err = c.conn.DomainCreateXML(string(domainXML), 0)
	return err
}
