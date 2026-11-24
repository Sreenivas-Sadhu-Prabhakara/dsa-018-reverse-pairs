package main

import "testing"

func TestReversePairs(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{1, 3, 2, 3, 1}, 2},
		{[]int{2, 4, 3, 5, 1}, 3},
		{[]int{5, 4, 3, 2, 1}, 4},
		{[]int{1, 2, 3}, 0},
		{[]int{}, 0},
		{[]int{-5, -5}, 1},
		{[]int{2147483647, 2147483647, -2147483647}, 2},
	}
	for _, c := range cases {
		if got := ReversePairs(c.in); got != c.want {
			t.Errorf("ReversePairs(%v) = %d, want %d", c.in, got, c.want)
		}
	}
}
