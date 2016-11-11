package ch0203

import (
	"testing"
	"reflect"
)

func TestMergeSort(t *testing.T) {
	 tcs := []struct{
		got []int
		want []int
		desc string
	}{
  		{[]int{1, 2, 3, 4, 2, 3, 4, 5}, []int{1, 2, 2, 3, 3, 4, 4,  5}, "typical case"},
  		{[]int{5, 4, 3, 3, 2, 2, 1}, []int{1, 2, 2, 3, 3, 4, 5}, "reversed"},
	}

	for _, tc := range tcs {
	  if got, want := mergeSort(tc.got), tc.want; !reflect.DeepEqual(got, want) {
	    t.Errorf("%s failed: got %v, want %v ", tc.desc, got, want)
	  }
	}
}
func TestMergeSortDuplicates(t *testing.T) {
	got := []int{5, 2, 4, 6, 1, 3, 1}
	want := []int{1, 1, 2, 3, 4, 5, 6}
	got = mergeSort(got)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("insertSort():  got %v, want %v", got, want)
	}
}

func TestMergeSortLargeValues(t *testing.T) {
	got := []int{5, 2, 2584, 6, 1, 3}
	want := []int{1, 2, 3, 5, 6, 2584}
	got = mergeSort(got)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("insertSort():  got %v, want %v", got, want)
	}
}

func TestMerge(t *testing.T){
	a := []int{1,5,29}
	b := []int{2,3,6} 
	want := []int{1,2,3,5,6,29}
	got := merge(a,b)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("insertSort():  got %v, want %v", got, want)
	}
}
func TestSplit(t *testing.T){
	a := []int{1,5,29,2,3,6}
	wantA := []int{1,5,29}
	wantB := []int{2,3,6}
	gotA, gotB := split(a)
	if !reflect.DeepEqual(wantA, gotA) {
		t.Errorf("insertSort():  got %v , want %v", gotA, wantA)
	}
	if !reflect.DeepEqual(wantB, gotB) {
		t.Errorf("insertSort():  got %v , want %v", gotB, wantB)
	}
}
