package ch02

import (
	"testing"
	"reflect"
)

func TestInsertSort(t *testing.T) {
	got := []int{5, 2, 4, 6, 1, 3}
	want := []int{1, 2, 3, 4, 5, 6}
	insertSort(got)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("insertSort():  got %v, want %v", got, want)
	}
}
func TestInsertSortShort(t *testing.T) {
	got := []int{5, 2}
	want := []int{2, 5}
	insertSort(got)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("insertSort():  got %v, want %v", got, want)
	}
}
func TestInsertSortDuplicates(t *testing.T) {
	got := []int{5, 2, 4, 6, 1, 3, 1}
	want := []int{1, 1, 2, 3, 4, 5, 6}
	insertSort(got)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("insertSort():  got %v, want %v", got, want)
	}
}

func TestInsertSortLargeValues(t *testing.T) {
	got := []int{5, 2, 2584, 6, 1, 3}
	want := []int{1, 2, 3, 5, 6, 2584}
	insertSort(got)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("insertSort():  got %v, want %v", got, want)
	}
}

func TestReverseInsertSort(t *testing.T) {
	got := []int{5, 2, 4, 6, 1, 3}
	want := []int{1, 2, 3, 4, 5, 6}
	reverseInsertSort(got)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("insertSort():  got %v, want %v", got, want)
	}
}
func TestReverseInsertSortshort(t *testing.T) {
	got := []int{5, 2}
	want := []int{2, 5}
	reverseInsertSort(got)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("insertSort():  got %v, want %v", got, want)
	}
}
func TestReverseInsertSortDuplicates(t *testing.T) {
	got := []int{5, 2, 4, 6, 1, 3, 1}
	want := []int{1, 1, 2, 3, 4, 5, 6}
	reverseInsertSort(got)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("insertSort():  got %v, want %v", got, want)
	}
}

func TestReverseInsertSortLargeValues(t *testing.T) {
	got := []int{5, 2, 2584, 6, 1, 3}
	want := []int{1, 2, 3, 5, 6, 2584}
	reverseInsertSort(got)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("insertSort():  got %v, want %v", got, want)
	}
}
