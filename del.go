package ipset

func (c *Config) Del(address string) error {
	args := []string{"del", c.Set, address}

	// logger
	c.Logger.Printf("Deleting entry '%s' from ipset '%s'", address, c.Set)

	return c.do(args)
}
