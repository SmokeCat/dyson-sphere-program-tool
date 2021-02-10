package util

import (
	"reflect"
	"sort"
	"testing"
)

func TestMergeMap(t *testing.T) {
	var (
		in1      = map[string]float64{"a": 2, "b": 4, "c": 6}
		in2      = map[string]float64{"a": 2, "c": 2.5}
		expected = map[string]float64{"a": 4, "b": 4, "c": 8.5}
	)

	MergeMap(in1, in2)

	if !reflect.DeepEqual(in1, expected) {
		t.Errorf("in1 = %v; expected = %v", in1, expected)
	}
}

func TestGetStringMapKeys(t *testing.T) {
	var (
		in       = map[string]interface{}{"a": "vala", "c": "valc", "b": 22}
		expected = []string{"a", "b", "c"}
	)

	actual := GetStringMapKeys(in)

	sort.Strings(actual)
	sort.Strings(expected)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v; expected = %v", actual, expected)
	}
}

func TestGetIntMapKeys(t *testing.T) {
	var (
		in       = map[int]interface{}{1: "val1", 3: "val3", 2: 22}
		expected = []int{1, 2, 3}
	)

	actual := GetIntMapKeys(in)

	sort.Ints(actual)
	sort.Ints(expected)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v; expected = %v", actual, expected)
	}
}
