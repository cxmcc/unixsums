package bsdsum

import (
	"fmt"
	"io"
	"testing"
)

type test struct {
	sum uint16
	in  string
}

var golden = []test{
	{97, "a"},
	{32914, "ab"},
	{16556, "abc"},
	{8378, "abcd"},
	{4290, "abcde"},
	{2247, "abcdef"},
	{33994, "abcdefg"},
	{17101, "abcdefgh"},
	{41423, "abcdefghi"},
	{53585, "abcdefghij"},
}

func TestGolden(t *testing.T) {
	for _, g := range golden {
		c := New()
		io.WriteString(c, g.in)
		s := c.Sum16()
		if s != g.sum {
			t.Errorf("Bsdsum(%s) = %d want %d", g.in, s, g.sum)
		}
	}
}

func ExampleNew() {
	c := New()
	io.WriteString(c, "Go is an open source programming environment.")
	fmt.Printf("sum: %d", c.Sum16())
	// Output: sum: 16973
}

func ExampleBsdsum() {
	data := []byte("Go is expressive, concise, clean, and efficient.")
	fmt.Printf("sum: %d", Bsdsum(data))
	// Output: sum: 8092
}
