package ipset

func (c *Config) Test(address string) bool {
	// logger
	c.Logger.Printf("Testind is entry '%s' in set '%s'", address, c.Set)

	args := []string{"test", c.Set, address}

	_, err := c.get(args)
	if err != nil {
		return false
	}

	return true
}
