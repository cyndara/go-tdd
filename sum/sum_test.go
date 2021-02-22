package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("testing for any size", func(t *testing.T) {
		numbers := []int{10, 20, 30}

		got := Sum(numbers)
		want := 60

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

}

func TestSumAllTails(t *testing.T) {

	assertCorrectSum := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d given", got, want)
		}
	}

	t.Run("testing for two slices", func(t *testing.T) {

		got := SumAllTails([]int{1, 2}, []int{3, 4, 5})
		want := []int{2, 9}
		assertCorrectSum(t, got, want)

	})

	t.Run("testing for empty slice", func(t *testing.T) {

		got := SumAllTails([]int{3, 4, 5}, []int{})
		want := []int{9, 0}
		assertCorrectSum(t, got, want)

	})

}
