package ipset

// Create new ipset.
func (c *Config) Create(options *CreateOptions) error {
	optionsArgs, err := options.GenArgs()
	if err != nil {
		return err
	}

	args := []string{"create", c.Set}

	// -exist
	if c.IgnoreExist {
		args = append(args, "-exist")
	}

	args = append(args, optionsArgs...)

	// logger
	c.Logger.Printf("Creating ipset '%s'", c.Set)

	return c.do(args)
}
