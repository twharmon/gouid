# Gouid

![](https://github.com/twharmon/gouid/workflows/Test/badge.svg) [![](https://goreportcard.com/badge/github.com/twharmon/gouid)](https://goreportcard.com/report/github.com/twharmon/gouid) [![](https://gocover.io/_badge/github.com/twharmon/gouid)](https://gocover.io/github.com/twharmon/gouid)

Use Gouid to create cryptographically secure random IDs. IDs can be byte slices or strings, both generated with just one allocation (see benchmarks below).
Strings IDs use a-z and 0-9 all lowercase.


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
	s := gouid.String(10)
	fmt.Println(s) // mpln6nq37p

	b := gouid.Bytes(16)
	fmt.Println(b) // [244 188 217 137 122 245 94 126 80 119 87 170 6 178 228 179]
}
```

## Benchmarks

```
BenchmarkString8 	        76.0 ns/op	       8 B/op	       1 allocs/op
BenchmarkString16	         101 ns/op	      16 B/op	       1 allocs/op
BenchmarkString32 	         143 ns/op	      32 B/op	       1 allocs/op
BenchmarkBytes8  	        64.7 ns/op	       8 B/op	       1 allocs/op
BenchmarkBytes16 	        80.9 ns/op	      16 B/op	       1 allocs/op
BenchmarkBytes32 	         121 ns/op	      32 B/op	       1 allocs/op
```

## Contribute

Make a pull request.
