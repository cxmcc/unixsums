package morehash

import "hash"

type Hash16 interface {
	hash.Hash
	Sum16() uint16
}
