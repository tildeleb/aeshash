package aeshash

import (
	"leb.io/hashland/nhash"
)

// nhash compatible interface
func NewAES(seed uint64) nhash.Hash64 {
	s := n(seed)
	return s
}

func (s *State) Hash64(b []byte, seeds ...uint64) uint64 {
	switch len(seeds) {
	case 1:
		s.seed = seeds[0]
	}
	s.hash = Hash(b, s.seed)
	//fmt.Printf("pc=0x%08x, pb=0x%08x\n", d.pc, d.pb)
	return s.hash
}

//func Hash(p unsafe.Pointer, s, h uintptr) uintptr
//func HashStr(p string, s, h uintptr) uintptr
//func aeshash(p unsafe.Pointer, s, h uintptr) uintptr
//func aeshash32(p unsafe.Pointer, s, h uintptr) uintptr
//func aeshash64(p unsafe.Pointer, s, h uintptr) uintptr
//func aeshashstr(p unsafe.Pointer, s, h uintptr) uintptr
