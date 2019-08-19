package utils

import "testing"

func TestUnique(t *testing.T) {
	arr := []string{}
	arr = Unique(arr)
	if len(arr) != 0 {
		t.Errorf(`Unique({}) = %v; want {}`, arr)
	}
	arr = []string{"a"}
	arr = Unique(arr)
	if len(arr) != 1 || arr[0] != "a" {
		t.Errorf(`Unique({"a"}) = %v; want {"a"}`, arr)
	}
	arr = []string{"a", "a"}
	arr = Unique(arr)
	if len(arr) != 1 || arr[0] != "a" {
		t.Errorf(`Unique({"a", "a"}) = %v; want {"a"}`, arr)
	}
	arr = []string{"a", "b", "c"}
	arr = Unique(arr)
	if len(arr) != 3 || arr[0] != "a" || arr[1] != "b" || arr[2] != "c" {
		t.Errorf(`Unique({"a", "b", "c"}) = %v; want {"a", "b", "c"}`, arr)
	}
	arr = []string{"a", "b", "a", "c", "c", "a"}
	arr = Unique(arr)
	if len(arr) != 3 || arr[0] != "a" || arr[1] != "b" || arr[2] != "c" {
		t.Errorf(`Unique({"a", "b", "a", "c", "c", "a"}) = %v; want {"a", "b", "c"}`, arr)
	}
}
