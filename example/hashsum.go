package main

import (
	"fmt"
	"hash"
	"hash/adler32"
	"hash/crc32"
	"hash/fnv"
	"io"
	"log"
	"os"

	"github.com/spaolacci/murmur3"
	"github.com/urfave/cli"
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

func main() {
	app := cli.NewApp()
	app.Name = "hashsum"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "algo",
			Value: "xxHash",
			Usage: "hash algorithm [xxHash|crc|crc-c|crc-k|adler|fnv|fnva|murmur]",
		}}
	app.Action = func(c *cli.Context) {
		hashfn := getHasher32(c.String("algo"))
		if hashfn == nil {
			log.Fatalln("invalid argument", c.String("algo"))
		}
		io.Copy(hashfn, os.Stdin)
		fmt.Printf("%x\n", hashfn.Sum(nil))
	}
	app.Run(os.Args)
}
