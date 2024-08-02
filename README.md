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
// ipset create Test hash:net family inet
func main() {
	// Create config: path to ipset bin and name of set.
	config, err := ipset.NewConfig("/usr/sbin/ipset", "Test")
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

### Flush ipset

```go
package main

import (
	"log"
	"os"

	"github.com/exsver/go-iptables"
)

// !!! Requires root privileges
//
// Create and fill new ipset
//   ipset create FlushTest hash:net family inet
//   ipset add FlushTest 10.10.10.10
//   ipset add FlushTest 192.168.1.0/24
func main() {
	// Create config: path to ipset bin and name of set.
	config, err := ipset.NewConfig("/usr/sbin/ipset", "FlushTest")
	if err != nil {
		log.Fatal(err)
	}

	// Optional: Set debug logger
	config.SetLogger(log.New(os.Stdout, "Debug: ", 0))

	// Flush ipset
	err := config.Flush()
	if err != nil {
		log.Fatal(err)
	}
}
```
