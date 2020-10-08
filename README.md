# Gouid

![](https://github.com/twharmon/gouid/workflows/Test/badge.svg) [![](https://goreportcard.com/badge/github.com/twharmon/gouid)](https://goreportcard.com/report/github.com/twharmon/gouid) [![](https://gocover.io/_badge/github.com/twharmon/gouid)](https://gocover.io/github.com/twharmon/gouid) [![GoDoc](https://godoc.org/github.com/twharmon/gouid?status.svg)](https://godoc.org/github.com/twharmon/gouid)

Use Gouid to create cryptographically secure random strings.


## Documentation

For full documentation see [pkg.go.dev](https://pkg.go.dev/github.com/twharmon/gouid).


## Example

```go
package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/twharmon/gouid"
)

func main() {
	uid := gouid.New(10)
	fmt.Println(uid) // mpln6nq37p
}
```

## Benchmarks

```
BenchmarkNew6    	17366401	        65.6 ns/op	       8 B/op	       1 allocs/op
BenchmarkNew32   	 9240732	       133 ns/op	      32 B/op	       1 allocs/op
```

## Contribute

Make a pull request.
