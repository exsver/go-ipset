package ipset

// Flush ipset.
func (c *Config) Flush() error {
	args := []string{"flush", c.Set}

	// logger
	c.Logger.Printf("Flushing ipset '%s'", c.Set)

	return c.do(args)
}
