package sorting

func ThreeWayQuickSort(arr []int) {
	// add random shuffle for better performace
	threeWayQuickSort(arr, 0, len(arr)-1)
}

// three way quick sorting algorithm is used for sorting array with value repeating
// variants are
// let 3 index be lt,i,gt
// [l0...lt] are smaller then partion key [v]
// [lt...i] are equal to v
// [gt...hi] are greater then v
func threeWayQuickSort(arr []int, lo, hi int) {
	if lo >= hi{
		return
	}
	i := lo
	lt := lo
	gt := hi
	v := arr[lo]
	for i <= gt {
		if arr[i] < v {
			arr[i], arr[lt] = arr[lt], arr[i]
			i++
			lt++
		} else if arr[i] > v {
			arr[i], arr[gt] = arr[gt], arr[i]
			gt--
		} else {
			i++
		}
	}
	threeWayQuickSort(arr, lo, lt-1)
	threeWayQuickSort(arr, gt+1, hi)
}
