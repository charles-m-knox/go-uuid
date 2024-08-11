package uuid

import (
	"crypto/rand"
	"encoding/hex"
	"io"
)

// New generates a uuid v4 variant 10. If an error is encountered when trying
// to acquire cryptographically secure randomness for the UUID, the app will
// panic.
func New() string {
	var uuid [16]byte

	_, err := io.ReadFull(rand.Reader, uuid[:])
	if err != nil {
		panic("uuid unable to get random bits")
	}

	uuid[6] = (uuid[6] & 0x0f) | 0x40 // v4
	uuid[8] = (uuid[8] & 0x3f) | 0x80 // variant 10

	var out [36]byte
	o := out[:]

	hex.Encode(o, uuid[:4])
	o[8] = '-'
	hex.Encode(o[9:13], uuid[4:6])
	o[13] = '-'
	hex.Encode(o[14:18], uuid[6:8])
	o[18] = '-'
	hex.Encode(o[19:23], uuid[8:10])
	o[23] = '-'
	hex.Encode(o[24:], uuid[10:])

	return string(o)
}
