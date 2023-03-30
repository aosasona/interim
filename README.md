# InterimDB

Interim is a fast in-memory key-value store for Golang.

# Installation

```bash
go get github.com/aosasona/interim
```

# Usage

```go
package main

import (
	"fmt"

	"github.com/aosasona/interim"
)

func main() {
	db := interim.New(interim.Config{
		CacheSize: 6, // optional, but depending on how much reads and write you'll be doing, you may want to tweak this (default is 8)
	})

	err := db.Set("some_key", "some_value") // Yes, you can use any type here
	if err != nil {
		fmt.Println(err)
	}

	var target string
	err = db.Get("some_key", &target) // Make sure your target is the same type as the data type you put in
	if err != nil || target == "" {
		fmt.Printf("something is broken: %v", err)
	}

	err = db.Delete("some_key")
	if err != nil {
		fmt.Println(err)
	}

	if !db.Exists("some_key") {
		fmt.Println("some_key not found!")
	}
}
```

See performance test results [here](https://github.com/aosasona/interim-perf-test).
