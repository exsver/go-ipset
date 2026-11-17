package ipset

// Rename ipset. On success, Config.Set is updated to the new name.
func (c *Config) Rename(newName string) error {
	args := []string{"rename", c.Set, newName}

	// logger
	c.Logger.Printf("Renaming ipset '%s' to '%s'", c.Set, newName)

	err := c.do(args)
	if err != nil {
		return err
	}

	c.Set = newName

	return nil
}
