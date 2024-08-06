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

func (o *CreateOptions) GenArgs() ([]string, error) {
	var args []string

	// Type
	switch o.Type {
	case "hash:ip", "hash:mac", "hash:ip,mac", "hash:net", "hash:net,net", "hash:ip,port", "hash:net,port", "hash:ip,port,ip", "hash:ip,port,net", "hash:ip,mark", "hash:net,port,net", "hash:net,iface", :
		args = append(args, o.Type)
		if o.MaxElem != 0 {
			args = append(args, "maxelem", strconv.Itoa(o.MaxElem))
		}
	case "list:set":
		args = append(args, o.Type)
	default:
		return nil, fmt.Errorf("unknown set type: %s", o.Type)
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
