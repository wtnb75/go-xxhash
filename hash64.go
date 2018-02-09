package xxhash

import (
	//#cgo LDFLAGS: -lxxhash
	//#include <xxhash.h>
	"C"
	"encoding/binary"
	"fmt"
	"runtime"
	"unsafe"
)

type XxHash64 struct {
	State C.struct_XXH64_state_s
}

func NewXxHash64(seed int) *XxHash64 {
	res := XxHash64{}
	res.State = *C.XXH64_createState()
	C.XXH64_reset(&res.State, C.ulonglong(seed))
	runtime.SetFinalizer(&res, func(a *XxHash64) {
		C.XXH64_freeState(&a.State)
	})
	return &res
}

func (x *XxHash64) Sum64() uint64 {
	return uint64(C.XXH64_digest(&x.State))
}

func (x *XxHash64) Write(p []byte) (n int, err error) {
	if errcode := C.XXH64_update(&x.State, unsafe.Pointer(&p[0]), C.size_t(len(p))); errcode != C.XXH_OK {
		return 0, fmt.Errorf("digest error")
	}
	return len(p), nil
}

func (x *XxHash64) Reset() {
	C.XXH64_reset(&x.State, 0)
}

func (x *XxHash64) Size() int {
	return 8
}

func (x *XxHash64) BlockSize() int {
	return 8 * 4
}

func (x *XxHash64) Sum(b []byte) []byte {
	var res uint64
	if b == nil {
		res = x.Sum64()
	} else {
		res = uint64(C.XXH64(unsafe.Pointer(&b[0]), C.size_t(len(b)), 0))
	}
	bs := make([]byte, 8)
	binary.BigEndian.PutUint64(bs, res)
	return bs
}
