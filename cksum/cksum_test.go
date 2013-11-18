package cksum

import (
	"io"
	"testing"
)

type test struct {
	sum uint32
	in  string
}

var golden = []test{
	{1220704766, "a"},
	{2072780115, "ab"},
	{1219131554, "abc"},
	{1278160200, "abcd"},
	{996742021, "abcde"},
	{773139377, "abcdef"},
	{2915707922, "abcdefg"},
	{1095960684, "abcdefgh"},
	{1133324850, "abcdefghi"},
	{966588298, "abcdefghij"},
}

func TestGolden(t *testing.T) {
	for _, g := range golden {
		c := New()
		io.WriteString(c, g.in)
		s := c.Sum32()
		if s != g.sum {
			t.Errorf("Cksum(%s) = %d want %d", g.in, s, g.sum)
		}
	}
}
