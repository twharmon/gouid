# Gouid
[![Go Reference](https://pkg.go.dev/badge/github.com/twharmon/gouid.svg)](https://pkg.go.dev/github.com/twharmon/gouid) ![](https://github.com/twharmon/gouid/workflows/Test/badge.svg) [![](https://goreportcard.com/badge/github.com/twharmon/gouid)](https://goreportcard.com/report/github.com/twharmon/gouid)  [![codecov](https://codecov.io/gh/twharmon/gouid/branch/main/graph/badge.svg?token=K0P59TPRAL)](https://codecov.io/gh/twharmon/gouid)

Use Gouid to create cryptographically secure random IDs. IDs can be byte slices or strings, both generated with just one allocation (see benchmarks below).


## Documentation

For full documentation see [pkg.go.dev](https://pkg.go.dev/github.com/twharmon/gouid).


## Example

```go
package main

import (
	"fmt"
	
	"github.com/twharmon/gouid"
)

func main() {
	a := gouid.String(8, gouid.Secure32Char)
	fmt.Println(a) // mp1nq34p

	b := gouid.String(16, gouid.Secure64Char)
	fmt.Println(b) // h-SoemLKa_QtoXgh

	c := gouid.Bytes(16)
	fmt.Println(c) // [244 188 217 137 122 245 94 126 80 119 87 170 6 178 228 179]
}
```

## Benchmarks

```
goos: linux
goarch: amd64
pkg: github.com/twharmon/gouid
cpu: AMD Ryzen 7 7840HS w/ Radeon 780M Graphics     
BenchmarkString8     	 337.6 ns/op	       8 B/op	       1 allocs/op
BenchmarkString16    	 359.1 ns/op	      16 B/op	       1 allocs/op
BenchmarkString32    	 363.5 ns/op	      32 B/op	       1 allocs/op
BenchmarkBytes8      	 327.0 ns/op	       8 B/op	       1 allocs/op
BenchmarkBytes16     	 334.0 ns/op	      16 B/op	       1 allocs/op
BenchmarkBytes32     	 334.7 ns/op	      32 B/op	       1 allocs/op
```

## Contribute

Make a pull request.
