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

type XxHash32 struct {
	State C.struct_XXH32_state_s
}

func NewXxHash32(seed int) *XxHash32 {
	res := XxHash32{}
	res.State = *C.XXH32_createState()
	C.XXH32_reset(&res.State, C.uint(seed))
	runtime.SetFinalizer(&res, func(a *XxHash32) {
		C.XXH32_freeState(&a.State)
	})
	return &res
}

func (x *XxHash32) Sum32() uint32 {
	return uint32(C.XXH32_digest(&x.State))
}

func (x *XxHash32) Write(p []byte) (n int, err error) {
	if errcode := C.XXH32_update(&x.State, unsafe.Pointer(&p[0]), C.size_t(len(p))); errcode != C.XXH_OK {
		return 0, fmt.Errorf("digest error")
	}
	return len(p), nil
}

func (x *XxHash32) Reset() {
	C.XXH32_reset(&x.State, 0)
}

func (x *XxHash32) Size() int {
	return 4
}

func (x *XxHash32) BlockSize() int {
	return 4 * 4
}

func (x *XxHash32) Sum(b []byte) []byte {
	var res uint32
	if b == nil {
		res = x.Sum32()
	} else {
		res = uint32(C.XXH32(unsafe.Pointer(&b[0]), C.size_t(len(b)), 0))
	}
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, res)
	return bs
}
