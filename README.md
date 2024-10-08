# go-ipset
Go bindings for ipset

## Examples
### Create new Ipset
```go
package main

import (
	"log"
	"os"

	"github.com/exsver/go-ipset"
)

// !!! Requires root privileges
func main() {
	// Create config: path to ipset bin and name of set.
	config, err := ipset.NewConfig("/usr/sbin/ipset", "CreateTest")
	if err != nil {
		log.Fatal(err)
	}

	// Optional: Set debug logger
	config.SetLogger(log.New(os.Stdout, "Debug: ", 0))

	// Prepare options
	createOptions := ipset.CreateOptions{
		Type:     "hash:net",
		Counters: false,
		Comment:  false,
	}

	// Exec ipset
	err = config.Create(&createOptions)
	if err != nil {
		log.Fatal(err)
	}
}
```

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
// ipset create AddTest hash:net family inet
func main() {
	// Create config: path to ipset bin and name of set.
	config, err := ipset.NewConfig("/usr/sbin/ipset", "AddTest")
	if err != nil {
		log.Fatal(err)
	}

	// Optional: Set debug logger
	config.SetLogger(log.New(os.Stdout, "Debug: ", 0))

	// Prepare entry
	entry := ipset.Entry{
		Address: "192.168.1.10/32",
	}

	// Exec ipset
	err = config.Add(&entry)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Delete entry from ipset

```go
package main

import (
	"log"
	"os"

	"github.com/exsver/go-ipset"
)

// !!! Requires root privileges
//
// Create and fill new ipset
//   ipset create DelTest hash:net family inet
//   ipset add DelTest 10.0.0.1
//   ipset add DelTest 10.0.1.0/24
func main() {
	// Create config: path to ipset bin and name of set.
	config, err := ipset.NewConfig("/usr/sbin/ipset", "DelTest")
	if err != nil {
		log.Fatal(err)
	}

	// Optional: Set debug logger
	config.SetLogger(log.New(os.Stdout, "Debug: ", 0))

	// Del entry from ipset
	err = config.Del("10.0.0.1")
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

	"github.com/exsver/go-ipset"
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
	err = config.Flush()
	if err != nil {
		log.Fatal(err)
	}
}
```
