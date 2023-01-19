package lifting

/*
#cgo CFLAGS: -I.
#include "addSeeds.h"
*/
import "C"
import (
	"log"
	"unsafe"
)

func addSeedsToList(structureSeed int64) []int64 {
	ret0 := int64(0)
	ret1 := int64(0)

	ret := int(C.addSeeds(C.int64_t(structureSeed), (*C.int64_t)(unsafe.Pointer(&ret0)), (*C.int64_t)(unsafe.Pointer(&ret1))))

	r := []int64{}
	if ret == 1 {
		r = []int64{int64(ret0)}
	} else if ret == 2 {
		r = []int64{int64(ret0), int64(ret1)}
	} else if ret == -1 {
		panic("Too many world seeds")
	}

	log.Println("Reversing structure seed ", structureSeed, " found ", r)

	return r
}

func StructureSeedToWorldSeeds(structureSeed int64) []int64 {
	// r :=
	// for i := 0; i < len(r); i++ {
	// 	s := r[i]
	// 	r[i] = nextLong(&s)
	// }
	// done in C
	return addSeedsToList(structureSeed)
}
