// Code generated by "stringer -type=typeKind"; DO NOT EDIT

package main

import "fmt"

const _typeKind_name = "kindInvalidkindIdentifierkindInt32kindUInt32kindInt64kindUInt64kindFloat32kindFloat64kindBoolkindArraykindSlicekindByteArraykindByteSlicekindStringkindOptional"

var _typeKind_index = [...]uint8{0, 11, 25, 34, 44, 53, 63, 74, 85, 93, 102, 111, 124, 137, 147, 159}

func (i typeKind) String() string {
	if i < 0 || i >= typeKind(len(_typeKind_index)-1) {
		return fmt.Sprintf("typeKind(%d)", i)
	}
	return _typeKind_name[_typeKind_index[i]:_typeKind_index[i+1]]
}
