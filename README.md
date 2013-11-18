morehash
========

Morehash is an implementation of popular hashing algorithms in Go Language.
(those are not included in official go packages)


Install the package:
~~~
go get github.com/cxmcc/morehash
~~~

Usage, via interface of [hash package](http://golang.org/pkg/hash/):
~~~ go
package main

import "github.com/cxmcc/morehash/cksum"
import "fmt"
import "io"

func main() {
  h := cksum.New()
  io.WriteString(h, "morehash")
  fmt.Printf("cksum: %d", h.Sum32())
}

~~~

Hash Algorithms Provided:
* bsdsum (sum, UNIXsum)
* cksum (UNIXcksum, POSIX cksum)
