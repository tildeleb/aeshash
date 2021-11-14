package aeshash

import (
	_ "unsafe"

	"leb.io/hashland/nhash"
)

const hashRandomBytes = 32 // used in asm_{386,amd64}.s
var masks [32]uint64
var shifts [32]uint64
var aeskeysched [hashRandomBytes]byte // this is really 2 x 128 bit round keys
var aesdebug [hashRandomBytes]byte

func aeshashbody()

// Use these functions for higest speed
func Hash(b []byte, seed uint64) uint64
func HashStr(s string, seed uint64) uint64
func Hash64(v uint64, s uint64) uint64
func Hash32(v uint32, s uint64) uint64

func init() {
	p := aeskeysched[:]
	p[0], p[1], p[2], p[3], p[4], p[5], p[6], p[7] = 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8
	p[8], p[9], p[10], p[11], p[12], p[13], p[14], p[15] = 0x9, 0xA, 0xB, 0xC, 0xD, 0xE, 0xF, 0x10
	p[16], p[17], p[18], p[19], p[20], p[21], p[22], p[23] = 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18
	p[24], p[25], p[26], p[27], p[28], p[29], p[30], p[31] = 0x19, 0x1A, 0x1B, 0x1C, 0x1D, 0x1E, 0x1F, 0xFF
	p = aesdebug[:]
	p[0], p[1], p[2], p[3], p[4], p[5], p[6], p[7] = 0xFF, 0, 0, 0, 0, 0, 0, 0xFE
	p[8], p[9], p[10], p[11], p[12], p[13], p[14], p[15] = 0xFD, 0, 0, 0, 0, 0, 0, 0xFC
}

//
type State struct {
	hash uint64
	seed uint64
	clen int
	tail []byte
}

func n(seed uint64) *State {
	s := new(State)
	s.seed = seed
	s.Reset()
	return s
}

// New returns state used for a aeshash.
func New(seed uint64) nhash.Hash64 {
	s := n(seed)
	return s
}

// Size returns the size of the resulting hash.
func (s *State) Size() int { return 8 }

// BlockSize returns the blocksize of the hash which in this case is 1 byte.
func (s *State) BlockSize() int { return 1 }

// NumSeedBytes returns the maximum number of seed bypes required. In this case 2 x 32
func (s *State) NumSeedBytes() int {
	return 8
}

// HashSizeInBits returns the number of bits the hash function outputs.
func (s *State) HashSizeInBits() int {
	return 64
}

// Reset the hash state.
func (s *State) Reset() {
	s.hash = 0
	s.clen = 0
	s.tail = nil
}

// Write accepts a byte stream p used for calculating the hash. For now this call is lazy and the actual hash calculations take place in Sum() and Sum32().
func (s *State) Write(p []byte) (nn int, err error) {
	l := len(p)
	s.clen += l
	s.tail = append(s.tail, p...)
	return l, nil
}

// Write64 accepts a uint64 stream p used for calculating the hash. For now this call is lazy and the actual hash calculations take place in Sum() and Sum32().
func (s *State) Write64(h uint64) (err error) {
	s.clen += 8
	s.tail = append(s.tail, byte(h>>56), byte(h>>48), byte(h>>40), byte(h>>32), byte(h>>24), byte(h>>16), byte(h>>8), byte(h))
	return nil
}

// Sum returns the current hash as a byte slice.
func (s *State) Sum(b []byte) []byte {
	s.hash = Hash(s.tail, s.seed)
	h := s.hash
	return append(b, byte(h>>56), byte(h>>48), byte(h>>40), byte(h>>32), byte(h>>24), byte(h>>16), byte(h>>8), byte(h))
}

// Sum64 returns the current hash as a 64 bit unsigned type.
func (s *State) Sum64() uint64 {
	s.hash = Hash(s.tail, s.seed)
	return s.hash
}
