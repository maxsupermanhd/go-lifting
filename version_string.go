// Code generated by "stringer -type=Version"; DO NOT EDIT.

package lifting

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[MC_1_8-0]
	_ = x[MC_1_9-1]
	_ = x[MC_1_10-2]
	_ = x[MC_1_11-3]
	_ = x[MC_1_12-4]
	_ = x[MC_1_13-5]
	_ = x[MC_1_13_1-6]
	_ = x[MC_1_14-7]
	_ = x[MC_1_15-8]
	_ = x[MC_1_16-9]
	_ = x[MC_1_16_1-10]
	_ = x[MC_1_17-11]
	_ = x[MC_1_18-12]
	_ = x[MC_1_19-13]
}

const _Version_name = "MC_1_8MC_1_9MC_1_10MC_1_11MC_1_12MC_1_13MC_1_13_1MC_1_14MC_1_15MC_1_16MC_1_16_1MC_1_17MC_1_18MC_1_19"

var _Version_index = [...]uint8{0, 6, 12, 19, 26, 33, 40, 49, 56, 63, 70, 79, 86, 93, 100}

func (i Version) String() string {
	if i < 0 || i >= Version(len(_Version_index)-1) {
		return "Version(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Version_name[_Version_index[i]:_Version_index[i+1]]
}