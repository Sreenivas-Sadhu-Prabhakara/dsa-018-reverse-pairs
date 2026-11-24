package main

import "fmt"

// ReversePairs counts pairs (i, j) with i < j and nums[i] > 2*nums[j],
// counted across the halves during merge sort. Input is not modified.
func ReversePairs(nums []int) int {
	buf := append([]int{}, nums...)
	tmp := make([]int, len(buf))
	return rpSort(buf, tmp)
}

func rpSort(a, tmp []int) int {
	n := len(a)
	if n < 2 {
		return 0
	}
	mid := n / 2
	count := rpSort(a[:mid], tmp[:mid]) + rpSort(a[mid:], tmp[mid:])

	// Both halves sorted: for each left element, advance a monotone pointer
	// over the right half while the pair condition holds.
	j := mid
	for i := 0; i < mid; i++ {
		for j < n && int64(a[i]) > 2*int64(a[j]) {
			j++
		}
		count += j - mid
	}

	i, j2, k := 0, mid, 0
	for i < mid && j2 < n {
		if a[i] <= a[j2] {
			tmp[k] = a[i]
			i++
		} else {
			tmp[k] = a[j2]
			j2++
		}
		k++
	}
	for i < mid {
		tmp[k] = a[i]
		i++
		k++
	}
	copy(a[:k], tmp[:k])
	return count
}

func main() {
	fmt.Println(ReversePairs([]int{1, 3, 2, 3, 1}))
}
