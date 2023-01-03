package lifting

//go:generate stringer -type=Version
type Version int

const (
	MC_1_8 Version = iota
	MC_1_9
	MC_1_10
	MC_1_11
	MC_1_12
	MC_1_13
	MC_1_13_1
	MC_1_14
	MC_1_15
	MC_1_16
	MC_1_16_1
	MC_1_17
	MC_1_18
	MC_1_19
	MC_NEWEST = MC_1_19
)

var (
	Versions = []Version{
		MC_NEWEST,
		MC_1_8,
		MC_1_9,
		MC_1_10,
		MC_1_11,
		MC_1_12,
		MC_1_13,
		MC_1_13_1,
		MC_1_14,
		MC_1_15,
		MC_1_16,
		MC_1_16_1,
		MC_1_17,
		MC_1_18,
		MC_1_19,
	}
)
