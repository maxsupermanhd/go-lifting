package main

// #cgo CFLAGS: -I${SRCDIR}/cubiomes/
// #cgo LDFLAGS: ${SRCDIR}/cubiomes/libcubiomes.a -lm
// #include "finders.h"
// #include "layers.h"
// #include "rng.h"
import "C"
import (
	"log"
	"unsafe"

	"github.com/davecgh/go-spew/spew"
)

type Datapoint struct {
	Salt             int32
	RegionX, RegionZ int32
	OffsetX, OffsetZ int32
	Offset           int32
}

func main() {
	log.Println("Hello world!")

	d := []Datapoint{
		NewDP(CB_Shipwreck, CB_MC_1_17, -4914, 86450),
		NewDP(CB_Shipwreck, CB_MC_1_17, 80990, 416),
		NewDP(CB_Shipwreck, CB_MC_1_17, -94661, 11342),
		NewDP(CB_Shipwreck, CB_MC_1_17, 27266, -102678),
		NewDP(CB_Shipwreck, CB_MC_1_17, -95387, -7280),
		NewDP(CB_Shipwreck, CB_MC_1_17, -97765, -117862),
		NewDP(CB_Shipwreck, CB_MC_1_17, 28243, -85390),
		NewDP(CB_Shipwreck, CB_MC_1_17, 15937, 24953),
		NewDP(CB_Shipwreck, CB_MC_1_17, 14317, 90564),
		NewDP(CB_Shipwreck, CB_MC_1_17, -48884, 8587),
		NewDP(CB_Swamp_Hut, CB_MC_1_17, 14356, 119044),
		NewDP(CB_Swamp_Hut, CB_MC_1_17, 65552, 55314),
		NewDP(CB_Swamp_Hut, CB_MC_1_17, -14444, 28386),
		NewDP(CB_Swamp_Hut, CB_MC_1_17, -60921, 157997),
		NewDP(CB_Swamp_Hut, CB_MC_1_17, -27804, -7483),
		NewDP(CB_Swamp_Hut, CB_MC_1_17, 101874, -102793),
		NewDP(CB_Swamp_Hut, CB_MC_1_17, 102981, -63215),
		NewDP(CB_Swamp_Hut, CB_MC_1_17, -16267, -80831),
		NewDP(CB_Swamp_Hut, CB_MC_1_17, -134646, -15922),
		NewDP(CB_Swamp_Hut, CB_MC_1_17, -66110, 92835),
	}
	spew.Dump(d)
	spew.Dump(crack(d))
}

func NewDP(st, mc C.int, cx, cz int) Datapoint {
	s := (*C.struct_StructureConfig)(C.malloc(C.size_t(unsafe.Sizeof(C.struct_StructureConfig{}))))
	b := C.getStructureConfig(st, mc, s)
	if b == 0 {
		log.Println("Wrong structure!")
	}
	spacing := int(s.regionSize)
	x := cx
	z := cz
	if cx < 0 {
		x = cx - spacing + 1
	}
	if cz < 0 {
		z = cz - spacing + 1
	}
	log.Println(spacing)
	d := Datapoint{
		Salt:    int32(s.salt),
		RegionX: int32(x / spacing),
		RegionZ: int32(z / spacing),
		OffsetX: int32(cx) - int32(x/spacing)*int32(spacing),
		OffsetZ: int32(cz) - int32(z/spacing)*int32(spacing),
		Offset:  int32(s.chunkRange),
	}
	C.free(unsafe.Pointer(s))
	return d
}

func crack(data []Datapoint) []int64 {
	structureSeeds := []int64{}
	// seen := 0
lowerLoop:
	for lower := int64(0); lower < int64(1)<<19; lower++ {
		if lower%((1<<19)/100) == 0 {
			log.Printf("%3.0f%% (%6d/%-6d) (so far %d)", float64(lower)/float64(1<<19)*100, lower, 1<<19, len(structureSeeds))
		}
		for i := 0; i < len(data); i++ {
			rngLower := setRegionSeed(lower, int64(data[i].RegionX), int64(data[i].RegionZ), int64(data[i].Salt))
			// seen++
			// if seen == 176 {
			// 	os.Exit(1)
			// }
			offsetX := data[i].OffsetX
			offsetZ := data[i].OffsetZ
			offset := data[i].Offset
			if nextInt(&rngLower, offset)%4 != offsetX%4 || nextInt(&rngLower, offset)%4 != offsetZ%4 {
				continue lowerLoop
			}
		}
	higherLoop:
		for higher := int64(0); higher < int64(1)<<(48-19); higher++ {
			seed := (higher << 19) | lower
			for i := 0; i < len(data); i++ {
				rngHigher := setRegionSeed(seed, int64(data[i].RegionX), int64(data[i].RegionZ), int64(data[i].Salt))
				offsetX := data[i].OffsetX
				offsetZ := data[i].OffsetZ
				offset := data[i].Offset
				if nextInt(&rngHigher, offset) != offsetX || nextInt(&rngHigher, offset) != offsetZ {
					continue higherLoop
				}
			}
			structureSeeds = append(structureSeeds, seed)
		}
	}
	return structureSeeds
}

const (
	RegionA = int64(341873128712)
	RegionB = int64(132897987541)
)

func setRegionSeed(worldSeed, rx, rz, salt int64) int64 {
	return (rx*RegionA + rz*RegionB + worldSeed + salt) ^ 25214903917
}

func next(seed *int64, bits int32) int32 {
	*seed = (*seed*0x5deece66d + 0xb) & ((1 << 48) - 1)
	return int32(int64(*seed) >> (48 - bits))
}

func nextInt(seed *int64, bound int32) int32 {
	if (bound & -bound) == bound {
		return int32(int64(bound) * int64(next(seed, 31)) >> 31)
	}
	bits := int32(0)
	val := int32(0)
	for {
		bits = next(seed, 31)
		val = bits % bound
		if bits-val+(bound-1) >= 0 {
			break
		}
	}
	return val
}
