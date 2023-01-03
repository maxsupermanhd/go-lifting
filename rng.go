package lifting

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

func nextLong(seed *int64) int64 {
	return (int64(next(seed, 32)) << 32) + int64(next(seed, 32))
}

// com.seedfinding.mccore.util.math.NextLongReverser.addSeedsToList
func addSeedsToList(structureSeed int64) []int64 {
	ret := []int64{}
	lowerBits := structureSeed & 0xffff_ffff
	upperBits := structureSeed >> 32
	if (lowerBits & 0x8000_0000) != 0 {
		upperBits += 1
	}
	bitsOfDanger := int64(1)
	lowMin := lowerBits<<16 - bitsOfDanger
	lowMax := ((lowerBits+1)<<16 - bitsOfDanger) - 1
	upperMin := ((upperBits << 16) - 107048004364969) >> bitsOfDanger
	m1lv := (lowMax*-33441+upperMin*17549)/(1<<31-bitsOfDanger) + 1
	m2lv := (lowMin*46603+upperMin*39761)/(1<<32-bitsOfDanger) + 1
	seed := (-39761*m1lv + 35098*m2lv)
	if (46603*m1lv+66882*m2lv)+107048004364969>>16 == upperBits {
		if seed>>16 == lowerBits {
			ret = append(ret, (254681119335897*seed+120305458776662)&0xffff_ffff_ffff)
		}
	}
	seed = (-39761*(m1lv+1) + 35098*m2lv)
	if (46603*(m1lv+1)+66882*m2lv)+107048004364969>>16 == upperBits {
		if seed>>16 == lowerBits {
			ret = append(ret, (254681119335897*seed+120305458776662)&0xffff_ffff_ffff)
		}
	}
	seed = (-39761*m1lv + 35098*(m2lv+1))
	if (46603*m1lv+66882*(m2lv+1))+107048004364969>>16 == upperBits {
		if seed>>16 == lowerBits {
			ret = append(ret, (254681119335897*seed+120305458776662)&0xffff_ffff_ffff)
		}
	}
	return ret
}

func StructureSeedToWorldSeeds(structureSeed int64) []int64 {
	r := addSeedsToList(structureSeed)
	for i := 0; i < len(r); i++ {
		s := r[i]
		r[i] = nextLong(&s)
	}
	return r
}
