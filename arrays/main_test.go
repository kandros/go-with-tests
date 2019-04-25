package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3}

	got := sum(numbers)
	want := 6

	if want != got {
		t.Errorf("got %d want %d received, %v", got, want, numbers)
	}
}

func TestSumAll(t *testing.T) {
	got := sumAll([]int{1, 2, 3}, []int{1, 2, 3, 4, 5})
	want := []int{6, 15}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := sumAllTails([]int{1, 2, 3}, []int{1, 2, 3, 4, 5})
		want := []int{5, 14}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := sumAllTails([]int{}, []int{1, 2, 3, 4, 5})
		want := []int{0, 14}
		checkSums(t, got, want)
	})
}
