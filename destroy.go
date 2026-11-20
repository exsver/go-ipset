package ipset

// Destroy ipset.
func (c *Config) Destroy() error {
	args := []string{"destroy", c.Set}

	// -exist
	if c.IgnoreExist {
		args = append(args, "-exist")
	}

	// logger
	c.Logger.Printf("Destroying ipset '%s'", c.Set)

	return c.do(args)
}
