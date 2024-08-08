package ipset

import (
	"encoding/xml"
	"fmt"
)

func List(c *Config) (*Ipset, error) {
	args := []string{"list", c.Set, "-output", "xml"}

	// logger
	c.Logger.Printf("List ipset '%s'", c.Set)

	data, err := c.get(args)
	if err != nil {
		return nil, err
	}

	var ipsets Ipsets

	err = xml.Unmarshal(data, &ipsets)
	if err != nil {
		return nil, err
	}

	if len(ipsets.Ipset) != 1 {
		return nil, fmt.Errorf("expected 1 ipset, got %d", len(ipsets.Ipset))
	}

	return &ipsets.Ipset[0], nil
}
