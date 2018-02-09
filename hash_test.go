package xxhash

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func check32String(data string, expected uint32, t *testing.T) {
	t.Log("check", data)
	xxh := NewXxHash32(0)
	exp := make([]byte, 4)
	binary.BigEndian.PutUint32(exp, expected)
	if r := xxh.Sum([]byte(data)); bytes.Compare(r, exp) != 0 {
		t.Error(data, "is not", expected, exp, "vs", r)
	}
	xxh.Write([]byte(data))
	if r := xxh.Sum32(); r != expected {
		t.Error(data, "is not", expected, r)
	}
	xxh.Write([]byte(data))
	if r := xxh.Sum32(); r == expected {
		t.Error(data, "x2 is", expected, r)
	}
	xxh.Reset()
	xxh.Write([]byte(data))
	if r := xxh.Sum32(); r != expected {
		t.Error(data, "is not", expected, r)
	}
}

func check64String(data string, expected uint64, t *testing.T) {
	t.Log("check", data)
	xxh := NewXxHash64(0)
	exp := make([]byte, 8)
	binary.BigEndian.PutUint64(exp, expected)
	if r := xxh.Sum([]byte(data)); bytes.Compare(r, exp) != 0 {
		t.Error(data, "is not", expected, exp, "vs", r)
	}
	xxh.Write([]byte(data))
	if r := xxh.Sum64(); r != expected {
		t.Error(data, "is not", expected, r)
	}
	xxh.Write([]byte(data))
	if r := xxh.Sum64(); r == expected {
		t.Error(data, "x2 is", expected, r)
	}
	xxh.Reset()
	xxh.Write([]byte(data))
	if r := xxh.Sum64(); r != expected {
		t.Error(data, "is not", expected, r)
	}
}

func Test32Hello(t *testing.T) {
	// echo -n 'hello' | xxh32sum -H0
	// fb0077f9  -
	// echo -n 'hello world' | xxh32sum -H0
	// cebb6622  -
	check32String("hello", 0xfb0077f9, t)
	check32String("hello world", 0xcebb6622, t)
}

func Test32Size(t *testing.T) {
	xxh := NewXxHash32(0)
	if xxh.Size() != 4 {
		t.Error("size is not 32 bit", xxh.Size())
	}
	if len(xxh.Sum([]byte("blabla"))) != xxh.Size() {
		t.Error("size is not Size()")
	}
}

func Test64Hello(t *testing.T) {
	// echo -n 'hello' | xxh64sum -H1
	// 26c7827d889f6da3  -
	// echo -n 'hello world' | xxh64sum -H1
	// 45ab6734b21e6968  -
	check64String("hello", 0x26c7827d889f6da3, t)
	check64String("hello world", 0x45ab6734b21e6968, t)
}

func Test64Size(t *testing.T) {
	xxh := NewXxHash64(0)
	if xxh.Size() != 8 {
		t.Error("size is not 64 bit", xxh.Size())
	}
	if len(xxh.Sum([]byte("blabla"))) != xxh.Size() {
		t.Error("size is not Size()")
	}
}
