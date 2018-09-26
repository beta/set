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

package set_test

import (
	"testing"

	"github.com/beta/set"
)

func TestNewInt64Set(t *testing.T) {
	cases := []struct {
		Values []int64
		Length int
	}{
		{[]int64{}, 0},
		{[]int64{1, 2, 3, 4, 5}, 5},
	}
	for no, c := range cases {
		set := set.NewInt64Set(c.Values...)
		if length := set.Len(); length != c.Length {
			t.Errorf("Case #%d, values: %v, expect length %d, get %d", no, c.Values, c.Length, length)
		}
	}
}

func TestSliceToInt64Set(t *testing.T) {
	people := []struct {
		ID   int64
		Name string
	}{
		{1, "Name 1"},
		{2, "Name 2"},
		{3, "Name 3"},
	}

	idSet := set.SliceToInt64Set(people, func(i int) int64 {
		return people[i].ID
	})

	if idSet.Len() != len(people) {
		t.Errorf("Wrong length, expect %d, get %d", len(people), idSet.Len())
	}
	for _, person := range people {
		if !idSet.Has(person.ID) {
			t.Errorf("Missing %d in set", person.ID)
		}
	}
}

func TestSetAndDelete(t *testing.T) {
	cases := []struct {
		IsSet  bool
		Value  int64
		Length int
	}{
		{IsSet: true, Value: 1, Length: 1},
		{IsSet: true, Value: 2, Length: 2},
		{IsSet: true, Value: 1, Length: 2},
		{IsSet: false, Value: 2, Length: 1},
		{IsSet: false, Value: 2, Length: 1},
		{IsSet: false, Value: 3, Length: 1},
		{IsSet: false, Value: 1, Length: 0},
	}

	set := set.NewInt64Set()
	if set.Len() != 0 {
		t.Errorf("Set should be empty")
		t.FailNow()
	}

	for no, c := range cases {
		if c.IsSet {
			set.Set(c.Value)
		} else {
			set.Delete(c.Value)
		}
		if length := set.Len(); length != c.Length {
			t.Errorf("Case #%d, isSet: %v, value: %v, expect length: %d, get: %d", no, c.IsSet, c.Value, c.Length, length)
		}
	}
}
