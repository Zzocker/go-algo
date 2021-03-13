// Package sorting is quick sort demonstration
package sorting

func QuickSort(arr []int) {
	// add random shuffle here
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, l0, hi int) {
	if l0 >= hi {
		return
	}
	j := partion(arr, l0, hi)
	quickSort(arr, l0, j-1)
	quickSort(arr, j+1, hi)
}

// l0 : partion index
// hi : last index
func partion(arr []int, l0, hi int) int {
	i := l0
	j := hi
	for j > i {

		for ; i < hi && arr[i] < arr[l0]; i++ {
		}
		for ; j > l0 && arr[j] > arr[l0]; j-- {
		}
		if i >= j {
			break
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	arr[j], arr[l0] = arr[l0], arr[j]
	return j
}
