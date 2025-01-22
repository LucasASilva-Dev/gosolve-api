package index

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	int1 := 10
	int2 := 20
	int3 := 30
	slice := []*int{&int1, &int2, &int3}
	slice2 := []*int{&int1, &int1, &int2}

	tests := []struct {
		name     string
		slice    []*int
		target   int
		expected int
		found    bool
	}{
		{"exact match", slice, 20, 1, true},
		{"no match, but closest value within 10% tolerance", slice, 21, 1, true},
		{"no match, and no value within 10% tolerance", slice, 50, -1, false},
		{"empty slice", []*int{}, 10, -1, false},
		{"slice with single element", []*int{&int1}, 10, 0, true},
		{"slice with duplicate elements", slice2, 10, 1, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualIndex, actualFound := binarySearch(tt.slice, tt.target)
			if actualIndex != tt.expected || actualFound != tt.found {
				t.Errorf("binarySearch(%v, %d) = (%d, %v), want (%d, %v)", tt.slice, tt.target, actualIndex, actualFound, tt.expected, tt.found)
			}
		})
	}
}

func TestBinarySearchPanics(t *testing.T) {
	// Ensure that the function does not panic when given a nil slice.
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("binarySearch(nil, 10) panicked, want no panic")
		}
	}()
	binarySearch(nil, 10)
}

func TestLookup(t *testing.T) {
	int1 := 10
	int2 := 20
	int3 := 30
	slice := []*int{&int1, &int2, &int3}

	tests := []struct {
		name      string
		index     []*int
		position  int
		wantIndex int
		wantFound bool
	}{
		{"nil IndexManager", nil, 10, 0, false},
		{"empty index", []*int{}, 10, -1, false},
		{"existing position", slice, 20, 1, true},
		{"non-existing position", slice, 25, -1, false},
		{"position at start of index", slice, 10, 0, true},
		{"position at end of index", slice, 30, 2, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var im *IndexManager
			if tt.index != nil {
				im = &IndexManager{indexMonitor: &indexMonitor{file: tt.index}}
			}
			gotIndex, gotFound := im.Lookup(tt.position)
			if gotIndex != tt.wantIndex || gotFound != tt.wantFound {
				t.Errorf("Lookup() = (%d, %v), want (%d, %v)", gotIndex, gotFound, tt.wantIndex, tt.wantFound)
			}
		})
	}
}
