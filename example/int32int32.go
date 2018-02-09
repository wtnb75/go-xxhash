package main

import (
	"encoding/binary"
	"flag"
	"hash"
	"hash/adler32"
	"hash/crc32"
	"hash/fnv"
	"log"
	"math"
	"time"

	"github.com/cheggaaa/pb"
	"github.com/golang-collections/go-datastructures/bitarray"
	"github.com/spaolacci/murmur3"
	"github.com/wtnb75/go-xxhash"
)

func getHasher32(h string) hash.Hash {
	switch h {
	case "xxHash":
		return xxhash.NewXxHash32(0)
	case "crc":
		return crc32.NewIEEE()
	case "crc-c":
		tbl := crc32.MakeTable(crc32.Castagnoli)
		return crc32.New(tbl)
	case "crc-k":
		tbl := crc32.MakeTable(crc32.Koopman)
		return crc32.New(tbl)
	case "adler":
		return adler32.New()
	case "fnv":
		return fnv.New32()
	case "fnva":
		return fnv.New32a()
	case "murmur":
		return murmur3.New32()
	default:
		return nil
	}
}

// check bijection (one-to-one onto mapping) uint32 -> uint32
func main() {
	flag.Parse()
	for _, algo := range flag.Args() {
		errors := 0
		dups := 0
		var i uint32
		hashfn := getHasher32(algo)
		ba := bitarray.NewBitArray(math.MaxUint32)
		data := make([]byte, 4)
		bar := pb.StartNew(math.MaxUint32)
		start := time.Now()
		for i = 0; i < math.MaxUint32; i++ {
			bar.Increment()
			binary.LittleEndian.PutUint32(data, i)
			res := hashfn.Sum(data)
			resi := binary.LittleEndian.Uint32(res)
			b, err := ba.GetBit(uint64(resi))
			if err != nil {
				log.Println("bitarray error", err)
				errors++
			}
			if b {
				log.Println("already set", i, resi)
				dups++
			}
			ba.SetBit(uint64(resi))
		}
		bar.FinishPrint("finished")
		log.Println("algo", algo, "errors", errors, "dups", dups, "elapsed", time.Now().Sub(start))
	}
}
