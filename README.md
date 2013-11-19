unixsums
========

Unixsums is an implementation of the legacy checksum utils (cksum, sum) in Go Language.

Install the package:
~~~
go get github.com/cxmcc/unixsums
~~~

Usage, via interface of [hash package](http://golang.org/pkg/hash/):
~~~ go
package main

import "github.com/cxmcc/unixsums/cksum"
import "fmt"
import "io"

func main() {
  h := cksum.New()
  io.WriteString(h, "unixsums")
  fmt.Printf("cksum: %d", h.Sum32())
}

~~~

Hash algorithms provided:
* bsdsum (BSD checksum, sum, UNIXsum)
* cksum (UNIX cksum, POSIX cksum)
* sysvsum (UNIX SystemV sum, SYSV checksum, sum -s)
