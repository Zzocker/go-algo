package sorting

// MergeSort : sort given array using merge sort
func MergeSort(arr []int) {
	aux := make([]int, len(arr))
	mSort(arr, aux, 0, len(arr)-1)
}

// BottomUpMergeSort : sorting without using recursion
func BottomUpMergeSort(arr []int) {
	aux := make([]int, len(arr))
	for i := 1; i < len(arr); i += i {
		for j := 0; j < len(arr)-i; j += 2 * i {
			min := len(arr) - 1
			if j+(2*i)-1 < min {
				min = j + (2 * i) - 1
			}
			merge(arr, aux, j, j+i-1, min)
		}
	}
}

func mSort(arr []int, aux []int, l, h int) {
	if l >= h {
		return
	}
	mid := (l + h) / 2
	mSort(arr, aux, l, mid)
	mSort(arr, aux, mid+1, h)
	merge(arr, aux, l, mid, h)
}

func merge(arr []int, aux []int, l, mid, h int) {
	// copy the array from [l,h] into aux array
	for i := l; i <= h; i++ {
		aux[i] = arr[i]
	}
	var i int = l
	var j int = mid + 1
	for k := l; k <= h; k++ {
		if j > h {
			arr[k] = aux[i]
			i++
		} else if i > mid {
			arr[k] = aux[j]
			j++
		} else if aux[j] < aux[i] {
			arr[k] = aux[j]
			j++
		} else {
			arr[k] = aux[i]
			i++
		}
	}
}
