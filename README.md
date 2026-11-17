# go-ipset
Go bindings for ipset

## Examples
### Create a new Ipset
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

	// Optional: ignore exist
	config.SetIgnoreExist(true)

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

### Add Entry to the existing ipset

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

	// Optional: ignore exist
	config.SetIgnoreExist(true)

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

### Destroy ipset

```go
package main

import (
	"log"
	"os"

	"github.com/exsver/go-ipset"
)

// !!! Requires root privileges
//
// Create new ipset
//   ipset create DestroyTest hash:net family inet
func main() {
	// Create config: path to ipset bin and name of set.
	config, err := ipset.NewConfig("/usr/sbin/ipset", "DestroyTest")
	if err != nil {
		log.Fatal(err)
	}

	// Optional: Set debug logger
	config.SetLogger(log.New(os.Stdout, "Debug: ", 0))

	// Destroy ipset
	err = config.Destroy()
	if err != nil {
		log.Fatal(err)
	}
}
```

### Rename ipset

```go
package main

import (
	"log"
	"os"

	"github.com/exsver/go-ipset"
)

// !!! Requires root privileges
//
// Create new ipset
//   ipset create RenameTest hash:net family inet
func main() {
	// Create config: path to ipset bin and name of set.
	config, err := ipset.NewConfig("/usr/sbin/ipset", "RenameTest")
	if err != nil {
		log.Fatal(err)
	}

	// Optional: Set debug logger
	config.SetLogger(log.New(os.Stdout, "Debug: ", 0))

	// Rename ipset. On success, config.Set is updated to the new name.
	err = config.Rename("RenamedTest")
	if err != nil {
		log.Fatal(err)
	}
}
```

### Swap ipsets

```go
package main

import (
	"log"
	"os"

	"github.com/exsver/go-ipset"
)

// !!! Requires root privileges
//
// Create two ipsets of compatible types
//   ipset create SwapTest hash:net family inet
//   ipset create SwapTestTmp hash:net family inet
//   ipset add SwapTestTmp 10.0.0.1
func main() {
	// Create config: path to ipset bin and name of set.
	config, err := ipset.NewConfig("/usr/sbin/ipset", "SwapTest")
	if err != nil {
		log.Fatal(err)
	}

	// Optional: Set debug logger
	config.SetLogger(log.New(os.Stdout, "Debug: ", 0))

	// Swap the content of SwapTest and SwapTestTmp
	err = config.Swap("SwapTestTmp")
	if err != nil {
		log.Fatal(err)
	}
}
```

### List ipset

```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/exsver/go-ipset"
)

// !!! Requires root privileges
//
// Create and fill new ipset
//   ipset create ListTest hash:net family inet counters comment timeout 3600
//   ipset add ListTest 10.0.0.1 comment "net 10"
//   ipset add ListTest 192.168.1.0/24 comment "net 192"
func main() {
	// Create config: path to ipset bin and name of set.
	config, err := ipset.NewConfig("/usr/sbin/ipset", "ListTest")
	if err != nil {
		log.Fatal(err)
	}

	// Optional: Set debug logger
	config.SetLogger(log.New(os.Stdout, "Debug: ", 0))

	// List ipset
	set, err := config.List()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Name: %s, Type: %s, Entries: %d\n", set.Name, set.Type, set.Header.NumEntries)
	for _, entry := range set.Members.Members {
		fmt.Printf("Member: %s, Bytes: %d, Packets: %d, Timeout: %d, Comment: %s\n",
			entry.Elem,
			entry.Bytes,
			entry.Packets,
			entry.Timeout,
			entry.Comment,
		)
	}
}
```

### Test entry in ipset

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
//   ipset create TestTest hash:net family inet
//   ipset add TestTest 10.0.0.1
func main() {
	// Create config: path to ipset bin and name of set.
	config, err := ipset.NewConfig("/usr/sbin/ipset", "TestTest")
	if err != nil {
		log.Fatal(err)
	}

	// Optional: Set debug logger
	config.SetLogger(log.New(os.Stdout, "Debug: ", 0))

	// Test if entry exists in ipset
	if config.Test("10.0.0.1") {
		log.Println("entry exists")
	} else {
		log.Println("entry not found")
	}
}
```
