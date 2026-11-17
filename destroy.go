package ipset

// Destroy ipset.
func (c *Config) Destroy() error {
	args := []string{"destroy", c.Set}

	// logger
	c.Logger.Printf("Destroying ipset '%s'", c.Set)

	return c.do(args)
}
