// Copyright (c) 2018 Beta Kuang
//
// This software is provided 'as-is', without any express or implied
// warranty. In no event will the authors be held liable for any damages
// arising from the use of this software.
//
// Permission is granted to anyone to use this software for any purpose,
// including commercial applications, and to alter it and redistribute it
// freely, subject to the following restrictions:
//
// 1. The origin of this software must not be misrepresented; you must not
//    claim that you wrote the original software. If you use this software
//    in a product, an acknowledgment in the product documentation would be
//    appreciated but is not required.
// 2. Altered source versions must be plainly marked as such, and must not be
//    misrepresented as being the original software.
// 3. This notice may not be removed or altered from any source distribution.

package set

import (
	"reflect"
)

// I64Set represents an int64 set.
type I64Set map[int64]struct{}

// NewInt64Set returns a new int64 set with initial values.
func NewInt64Set(values ...int64) I64Set {
	set := make(map[int64]struct{}, len(values))
	for _, v := range values {
		set[v] = struct{}{}
	}
	return I64Set(set)
}

// SliceToInt64Set generates an int64 set from slice.
// Function picker returns an int64 value from the i-th place of the slice.
//
// This function panics if slice is not a slice.
func SliceToInt64Set(slice interface{}, picker func(i int) int64) I64Set {
	rv := reflect.ValueOf(slice)
	if rv.Type().Kind() != reflect.Slice {
		panic("set: param is not a slice")
	}

	length := rv.Len()
	set := make(map[int64]struct{}, length)
	for i := 0; i < length; i++ {
		set[picker(i)] = struct{}{}
	}

	return I64Set(set)
}

// Len returns the length of set. len() also works on I64Set.
func (set I64Set) Len() int {
	return len(set)
}

// Has returns whether v is in set.
func (set I64Set) Has(v int64) bool {
	_, exist := set[v]
	return exist
}

// Set writes v into set.
func (set I64Set) Set(v int64) {
	set[v] = struct{}{}
}

// Delete deletes v from set.
func (set I64Set) Delete(v int64) {
	delete(set, v)
}
