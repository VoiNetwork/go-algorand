// Code generated by "stringer -type=hashKind"; DO NOT EDIT.

package ledger

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[accountHK-0]
	_ = x[assetHK-1]
	_ = x[appHK-2]
	_ = x[kvHK-3]
}

const _hashKind_name = "accountHKassetHKappHKkvHK"

var _hashKind_index = [...]uint8{0, 9, 16, 21, 25}

func (i hashKind) String() string {
	if i >= hashKind(len(_hashKind_index)-1) {
		return "hashKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _hashKind_name[_hashKind_index[i]:_hashKind_index[i+1]]
}
