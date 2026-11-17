package ipset

// Swap the content of ipset with another existing ipset.
func (c *Config) Swap(otherSet string) error {
	args := []string{"swap", c.Set, otherSet}

	// logger
	c.Logger.Printf("Swapping ipset '%s' with '%s'", c.Set, otherSet)

	return c.do(args)
}
