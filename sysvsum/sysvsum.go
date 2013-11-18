package sysvsum

import "github.com/cxmcc/morehash"

const Size = 2

type digest struct {
	sum uint32
}

func New() morehash.Hash16 {
	return &digest{0}
}

func (d *digest) Reset() {
	d.sum = 0
}

func (d *digest) Size() int {
	return Size
}

func (d *digest) BlockSize() int {
	return 1
}

func (d *digest) Write(data []byte) (int, error) {
	for _, b := range data {
		d.sum += uint32(b)
	}
	return len(data), nil
}

func (d *digest) Sum16() uint16 {
	r := (d.sum & 0xffff) + (d.sum >> 16)
	return uint16(r&0xffff + r>>16)
}

func (d *digest) Sum(in []byte) []byte {
	s := d.Sum16()
	return append(in, byte(s>>8), byte(s))
}

func Sysvsum(data []byte) uint16 {
	d := New()
	d.Write(data)
	return d.Sum16()
}
