package lifting

import (
	"log"
)

type Datapoint struct {
	Salt             int32
	RegionX, RegionZ int32
	OffsetX, OffsetZ int32
	Offset           int32
}

func ExampleLiftStructures() {
	d := []Datapoint{
		NewDP(Shipwreck, MC_1_17, -4914, 86450),
		NewDP(Shipwreck, MC_1_17, 80990, 416),
		NewDP(Shipwreck, MC_1_17, -94661, 11342),
		NewDP(Shipwreck, MC_1_17, 27266, -102678),
		NewDP(Shipwreck, MC_1_17, -95387, -7280),
		NewDP(Shipwreck, MC_1_17, -97765, -117862),
		NewDP(Shipwreck, MC_1_17, 28243, -85390),
		NewDP(Shipwreck, MC_1_17, 15937, 24953),
		NewDP(Shipwreck, MC_1_17, 14317, 90564),
		NewDP(Shipwreck, MC_1_17, -48884, 8587),
		NewDP(SwampHut, MC_1_17, 14356, 119044),
		NewDP(SwampHut, MC_1_17, 65552, 55314),
		NewDP(SwampHut, MC_1_17, -14444, 28386),
		NewDP(SwampHut, MC_1_17, -60921, 157997),
		NewDP(SwampHut, MC_1_17, -27804, -7483),
		NewDP(SwampHut, MC_1_17, 101874, -102793),
		NewDP(SwampHut, MC_1_17, 102981, -63215),
		NewDP(SwampHut, MC_1_17, -16267, -80831),
		NewDP(SwampHut, MC_1_17, -134646, -15922),
		NewDP(SwampHut, MC_1_17, -66110, 92835),
	}
	structureSeeds := LiftStructures(d)
	log.Printf("Got %d structure seeds: %v", len(structureSeeds), structureSeeds)
	worldSeeds := []int64{}
	for _, v := range structureSeeds {
		worldSeeds = append(worldSeeds, structureSeedToWorldSeeds(v)...)
	}
	log.Printf("Got %d world seeds: %v", len(worldSeeds), worldSeeds)
}

func NewDP(st Structure, mc Version, cx, cz int) Datapoint {
	s := getStructureConfig(st, mc)
	if s == nil {
		log.Fatal("Wrong structure!")
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
	d := Datapoint{
		Salt:    int32(s.salt),
		RegionX: int32(x / spacing),
		RegionZ: int32(z / spacing),
		OffsetX: int32(cx) - int32(x/spacing)*int32(spacing),
		OffsetZ: int32(cz) - int32(z/spacing)*int32(spacing),
		Offset:  int32(s.chunkRange),
	}
	return d
}

func LiftStructures(data []Datapoint) []int64 {
	structureSeeds := []int64{}
lowerLoop:
	for lower := int64(0); lower < int64(1)<<19; lower++ {
		if lower%((1<<19)/100) == 0 {
			log.Printf("lower %3.0f%% (%6d/%-6d) (so far %d)", float64(lower)/float64(1<<19)*100, lower, 1<<19, len(structureSeeds))
		}
		for i := 0; i < len(data); i++ {
			rngLower := setRegionSeed(lower, int64(data[i].RegionX), int64(data[i].RegionZ), int64(data[i].Salt))
			offsetX := data[i].OffsetX
			offsetZ := data[i].OffsetZ
			offset := data[i].Offset
			if nextInt(&rngLower, offset)%4 != offsetX%4 || nextInt(&rngLower, offset)%4 != offsetZ%4 {
				continue lowerLoop
			}
		}
	higherLoop:
		for higher := int64(0); higher < int64(1)<<(48-19); higher++ {
			if higher%((int64(1)<<(48-19))/100) == 0 {
				log.Printf("higher %3.0f%% (%6d/%-6d) (so far %d)", float64(higher)/float64(int64(1)<<(48-19))*100, higher, int64(1)<<(48-19), len(structureSeeds))
			}
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
