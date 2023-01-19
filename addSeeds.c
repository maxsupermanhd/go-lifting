#include "addSeeds.h"

#include <stdio.h>

long floor_div(long x, long y) {
	long r = x / y;
	if ((x ^ y) < 0 && (r * y != x)) {
		r-=1;
	}
	return r;
}

int next(uint64_t *seed, const int bits) {
	*seed = (*seed * 0x5deece66d + 0xb) & ((1ULL << 48) - 1);
	return (int) ((int64_t)*seed >> (48 - bits));
}

uint64_t nextLong(uint64_t *seed) {
	return ((uint64_t) next(seed, 32) << 32) + next(seed, 32);
}


int addSeeds(int64_t structureSeed, int64_t* ret0, int64_t* ret1) {
	int64_t lower_bits = structureSeed & 0xffffffffULL;
	int64_t upper_bits = (uint64_t)structureSeed >> 32;
	if ((lower_bits & 0x80000000ULL) != 0) {
		upper_bits += 1;
	}
	int64_t bits_of_danger = 1;
	int64_t low_min = lower_bits << 16 - bits_of_danger;
	int64_t low_max = ((lower_bits + 1) << 16 - bits_of_danger) - 1;
	int64_t upper_min = ((upper_bits << 16) - 107048004364969LL) >> bits_of_danger;
	int64_t m1lv = floor_div(low_max * -33441 + upper_min * 17549, 1LL << 31 - bits_of_danger) + 1;
	int64_t m2lv = floor_div(low_min * 46603 + upper_min * 39761, 1LL << 32 - bits_of_danger) + 1;
	int ret = 0;
	if (((uint64_t)(((46603 * m1lv + 66882 * m2lv) + 107048004364969LL)) >> 16) == upper_bits) {
		int64_t seed = -39761 * m1lv + 35098 * m2lv;
		if ((((uint64_t)seed) >> 16) == lower_bits) {
			uint64_t s = (254681119335897LL * seed + 120305458776662LL) & 0xffffffffffffLL;
			*ret0 = nextLong(&s);
			ret++;
		}
	}
	if (((uint64_t)((46603 * (m1lv+1) + 66882 * m2lv) + 107048004364969LL) >> 16) == upper_bits) {
		int64_t seed = -39761 * (m1lv+1)+ 35098 * m2lv;
		if ((((uint64_t)seed) >> 16) == lower_bits) {
			uint64_t s = (254681119335897LL * seed + 120305458776662LL) & 0xffffffffffffLL;
			if (ret == 0) {
				*ret0 = nextLong(&s);
			} else if (ret == 1) {
				*ret1 = nextLong(&s);
			}
			ret++;
		}
	}
	if (((uint64_t)((46603 * m1lv + 66882 * (m2lv+1)) + 107048004364969LL) >> 16) == upper_bits) {
		int64_t seed = -39761 * m1lv + 35098 * (m2lv +1);
		if ((((uint64_t)seed) >> 16) == lower_bits) {
			uint64_t s = (254681119335897LL * seed + 120305458776662LL) & 0xffffffffffffLL;
			if (ret == 0) {
				*ret0 = nextLong(&s);
			} else if (ret == 1) {
				*ret1 = nextLong(&s);
			}
			ret++;
		}
	}
	if (((uint64_t)((46603 * (m1lv+1) + 66882 * (m2lv+1)) + 107048004364969LL) >> 16) == upper_bits) {
		int64_t seed = -39761 * (m1lv+1)+ 35098 * (m2lv +1);
		if ((((uint64_t)seed) >> 16) == lower_bits) {
			uint64_t s = (254681119335897LL * seed + 120305458776662LL) & 0xffffffffffffLL;
			if (ret == 0) {
				*ret0 = nextLong(&s);
			} else if (ret == 1) {
				*ret1 = nextLong(&s);
			}
			ret++;
		}
	}
	return ret;
}