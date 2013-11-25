Unixsums is an implementation of the legacy checksum utils (cksum, sum) in Go Language.

-----

[![Build Status](https://drone.io/github.com/cxmcc/unixsums/status.png)](https://drone.io/github.com/cxmcc/unixsums/latest)
[![GoDoc](http://godoc.org/github.com/cxmcc/unixsums?status.png)](http://godoc.org/github.com/cxmcc/unixsums)

### Checksum functions provided
* bsdsum (BSD checksum, sum, UNIXsum)
* cksum (UNIX cksum, POSIX cksum)
* sysvsum (UNIX SystemV sum, SYSV checksum, sum -s)

### API Documentation
Implementing [hash.Hash](http://golang.org/pkg/hash/#Hash).  
Documentation currently available at Godoc: [http://godoc.org/github.com/cxmcc/unixsums](http://godoc.org/github.com/cxmcc/unixsums)

### Installing
~~~
go get github.com/cxmcc/unixsums
~~~

### Example
~~~ go
package main

import "github.com/cxmcc/unixsums/cksum"
import "fmt"
import "io"

func main() {
  h := cksum.New()
  io.WriteString(h, "Go is expressive, concise, clean, and efficient.")
  fmt.Printf("cksum: %d", h.Sum32())
  // Output: cksum: 1937373249
}
~~~
### License
It's MIT License
