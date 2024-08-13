package ipset

import "strings"

// Add entry to ipset.
func (c *Config) Add(entry *Entry) error {
	entryArgs, err := entry.GenArgs()
	if err != nil {
		return err
	}

	args := []string{"add", c.Set}

	// -exist
	if c.IgnoreExist {
		args = append(args, "-exist")
	}

	args = append(args, entryArgs...)

	// logger
	c.Logger.Printf("Adding entry '%s' to ipset '%s'", strings.Join(entryArgs, " "), c.Set)

	return c.do(args)
}
