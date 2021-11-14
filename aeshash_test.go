package aeshash

import (
	"testing"
)

func TestUnseededHash(t *testing.T) {
	m := map[uint64]struct{}{}
	for i := 0; i < 1000; i++ {
		h := NewAES(uint64(i))
		m[h.Sum64()] = struct{}{}
	}
	if len(m) < 900 {
		t.Errorf("empty hash not sufficiently random: got %d, want 1000", len(m))
	}
}

func TestSeededHash(t *testing.T) {
	m := map[uint64]struct{}{}
	for i := 0; i < 1000; i++ {
		h := NewAES(1)
		m[h.Sum64()] = struct{}{}
	}
	if len(m) != 1 {
		t.Errorf("seeded hash is random: got %d, want 1", len(m))
	}
}

/*
func TestFirstTenHashes(t *testing.T) {
var firstTenHashes = []uint64{
	0x54de6ee3c89ad535,
	0x54de6ee3c89ad535,
	0xd569b4f2ca17ba57,
	0x02e52663b7589218,
	0xbaba7ca05ac2c5e8,
	0xd2446078b660d417,
	0x89485dc9a8a45ce3,
	0xd72a00044a27326f,
	0x4af060bd80b05767,
	0xb457cd9b9b95cdbf,
	0x8782428a7c5cd1be,
}
*/

func TestHash64(t *testing.T) {
	m := map[uint64]struct{}{}
	for i := 0; i < 1000; i++ {
		h := NewAES(1)
		m[h.Sum64()] = struct{}{}
	}
	if len(m) != 1 {
		t.Errorf("seeded hash is random: got %d, want 1", len(m))
	}
}

func benchmarkStdSize(b *testing.B, size int) {
	h := NewAES(1)
	buf := make([]byte, size)
	b.SetBytes(int64(size))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		h.Reset()
		h.Write(buf)
		h.Sum64()
	}
}

func benchmarkSize(b *testing.B, size int) {
	buf := make([]byte, size)
	b.SetBytes(int64(size))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Hash(buf, uint64(i))
	}
}

func BenchmarkHash8Bytes(b *testing.B) {
	benchmarkSize(b, 8)
}

func BenchmarkHash320Bytes(b *testing.B) {
	benchmarkSize(b, 320)
}

func BenchmarkHash1K(b *testing.B) {
	benchmarkSize(b, 1024)
}

func BenchmarkHash8K(b *testing.B) {
	benchmarkSize(b, 8192)
}

func BenchmarkHash32(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Hash32(uint32(i), uint64(0))
	}
}

func BenchmarkHash64(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Hash64(uint64(i), uint64(0))
	}
}
