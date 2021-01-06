# Gouid

![](https://github.com/twharmon/gouid/workflows/Test/badge.svg) [![](https://goreportcard.com/badge/github.com/twharmon/gouid)](https://goreportcard.com/report/github.com/twharmon/gouid) [![](https://gocover.io/_badge/github.com/twharmon/gouid)](https://gocover.io/github.com/twharmon/gouid)

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
	s := gouid.String(8, gouid.LowerCaseAlphaNum)
	fmt.Println(s) // mp6nq37p

	s := gouid.String(16, gouid.MixedCaseAlpha)
	fmt.Println(s) // hCSoemLKaUQtoXgh

	b := gouid.Bytes(16)
	fmt.Println(b) // [244 188 217 137 122 245 94 126 80 119 87 170 6 178 228 179]
}
```

## Benchmarks

```
BenchmarkString8	 120 ns/op	       8 B/op	       1 allocs/op
BenchmarkString16	 197 ns/op	      16 B/op	       1 allocs/op
BenchmarkString32	 345 ns/op	      32 B/op	       1 allocs/op
BenchmarkBytes8	     67.3 ns/op	       8 B/op	       1 allocs/op
BenchmarkBytes16	 94.4 ns/op	      16 B/op	       1 allocs/op
BenchmarkBytes32	 143 ns/op	      32 B/op	       1 allocs/op
```

## Contribute

Make a pull request.
