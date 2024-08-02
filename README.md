# go-ipset
Go bindings for ipset

## Examples
### Add Entry to existing ipset

```go
package main

import (
	"log"
	"os"

	"github.com/exsver/go-ipset"
)

// !!! Requires root privileges
//
// Create ipset
// ipset create test hash:net family inet
func main() {
	// Create config: path to ipset bin and name of set.
	config, err := ipset.NewConfig("/usr/sbin/ipset", "test")
	if err != nil {
		log.Fatal(err)
	}

	// Optional: Set debug logger
	config.SetLogger(log.New(os.Stdout, "Debug: ", 0))

	// Prepare entry
	entry := ipset.Entry{
		Address: "192.168.1.10/32",
	}

	// Exec iptables
	err = config.Add(&entry)
	if err != nil {
		log.Fatal(err)
	}
}
```
