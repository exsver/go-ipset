package ipset

import (
	"fmt"
	"strconv"
)

type CreateOptions struct {
	// Type - Set type:
	// "hash:ip", "hash:net", "hash:ip,port", etc
	Type string
	// Timeout - the default timeout value (in seconds)
	Timeout int
	// Counters - enables counters
	Counters bool
	// Comment - enables comments
	Comment bool
	// MaxElem - the maximal number of elements which can be stored in the set
	MaxElem int
}

// maxTimeout - the maximal timeout value (in seconds) accepted by the kernel
const maxTimeout = 2147483

var bitmapTypes = map[string]bool{
	"bitmap:ip":     true,
	"bitmap:ip,mac": true,
	"bitmap:port":   true,
}

var hashTypes = map[string]bool{
	"hash:ip":           true,
	"hash:mac":          true,
	"hash:ip,mac":       true,
	"hash:net":          true,
	"hash:net,net":      true,
	"hash:ip,port":      true,
	"hash:net,port":     true,
	"hash:ip,port,ip":   true,
	"hash:ip,port,net":  true,
	"hash:ip,mark":      true,
	"hash:net,port,net": true,
	"hash:net,iface":    true,
}

var listTypes = map[string]bool{
	"list:set": true,
}

func (o *CreateOptions) Validate() error {
	// Type is required
	if o.Type == "" {
		return fmt.Errorf("set type is required")
	}

	// Validate type
	if !bitmapTypes[o.Type] && !hashTypes[o.Type] && !listTypes[o.Type] {
		return fmt.Errorf("unknown set type: %s", o.Type)
	}

	// Validate timeout
	if o.Timeout < 0 || o.Timeout > maxTimeout {
		return fmt.Errorf("invalid timeout %d: must be between 0 and %d", o.Timeout, maxTimeout)
	}

	// Validate maxelem
	if o.MaxElem < 0 {
		return fmt.Errorf("invalid maxelem %d: must not be negative", o.MaxElem)
	}

	return nil
}

func (o *CreateOptions) GenArgs() ([]string, error) {
	if err := o.Validate(); err != nil {
		return nil, err
	}

	// Type
	args := []string{o.Type}
	if hashTypes[o.Type] && o.MaxElem != 0 {
		args = append(args, "maxelem", strconv.Itoa(o.MaxElem))
	}

	// Timeout
	if o.Timeout != 0 {
		args = append(args, "timeout", strconv.Itoa(o.Timeout))
	}

	// Counters
	if o.Counters {
		args = append(args, "counters")
	}

	// Comment
	if o.Comment {
		args = append(args, "comment")
	}

	return args, nil
}
