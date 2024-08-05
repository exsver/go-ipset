package ipset

import (
	"fmt"
	"math"
	"strconv"
)

type Entry struct {
	Address string
	Comment string
	Timeout int
}

func (e *Entry) GenArgs() ([]string, error) {
	if !isStringValidCidrOrIP(e.Address) {
		return nil, fmt.Errorf("invalid address")
	}

	var args []string

	args = append(args, e.Address)

	if e.Timeout != 0 {
		args = append(args, "timeout", strconv.Itoa(int(math.Min(float64(e.Timeout), TimeoutMax))))
	}

	if e.Comment != "" {
		args = append(args, "comment", e.Comment)
	}

	return args, nil
}
