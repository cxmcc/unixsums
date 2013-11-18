package sysvsum

import (
	"io"
	"testing"
)

type test struct {
	sum uint16
	in  string
}

var golden = []test{
	{97, "a"},
	{195, "ab"},
	{294, "abc"},
	{394, "abcd"},
	{495, "abcde"},
	{597, "abcdef"},
	{700, "abcdefg"},
	{804, "abcdefgh"},
	{909, "abcdefghi"},
	{1015, "abcdefghij"},
}

func TestGolden(t *testing.T) {
	for _, g := range golden {
		c := New()
		io.WriteString(c, g.in)
		s := c.Sum16()
		if s != g.sum {
			t.Errorf("Sysvsum(%s) = %d want %d", g.in, s, g.sum)
		}
	}
}
