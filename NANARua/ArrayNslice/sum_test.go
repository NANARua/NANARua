package ArrayNslice

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collections of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		if want != got {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	// t.Run("collections of any size", func(t *testing.T) {
	// 	numbers := []int{1, 2, 3}
	// 	got := Sum(numbers)
	// 	want := 6
	// 	if want != got {
	// 		t.Errorf("got %d want %d given, %v", got, want, numbers)
	// 	}
	// })
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{3, 4, 5})
	want := 15
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d want %d", got, want)
	}
}
