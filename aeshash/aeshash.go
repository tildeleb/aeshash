package main

import (
	"flag"
	"fmt"
	"os"
	_ "unsafe"

	"leb.io/aeshash"
)

const blockSize = 1024 * 1024

func readFullAESHash(path string) (r uint64) {
	//fmt.Printf("readFullAESHash: file=%q\n", path)
	f, err := os.Open(path)
	if err != nil {
		fmt.Printf("file %q, err=%v", path, err)
	}
	defer f.Close()

	buf := make([]byte, blockSize)
	hf := aeshash.NewAES(0)
	hf.Reset()
	for {
		l, err := f.Read(buf)
		//fmt.Printf("f=%q, err=%v, l=%d, size=%d\n", fi.Name(), err, l, fi.Size())
		if l == 0 {
			break
		}
		if l < 0 || err != nil {
			fmt.Printf("file %q, err=%v", path, err)
			return
		}
		hf.Write(buf[:l])
	}
	r = hf.Sum64()
	//fmt.Printf("readFullHash: p=%q, r=%#016x\n", p, r)
	//h.Write(buf[0:l])
	//r = h.Sum64()
	//fmt.Printf("readFullHash: file=%q, hash=0x%016x\n", p, r)
	return r
}

func main() {
	var arg = flag.Uint64("i", 0, "number to hash")
	var seed = flag.Uint64("s", 0, "seed to hash")
	var zero = flag.Bool("z", false, "hash of 0")

	flag.Parse()
	switch {
	case *arg != 0 || *zero:
		fmt.Printf("%#016x\n", aeshash.Hash64(*arg, *seed))
	default:
		if len(flag.Args()) <= 0 {
			return
		}
		//fmt.Printf("main: nargs=%d\n", len(flag.Args()))
		for _, path := range flag.Args() {
			h := readFullAESHash(path)
			fmt.Printf("%016x\t%s\n", h, path)
		}
	}
}
