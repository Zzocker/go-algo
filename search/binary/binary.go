package binary

// LoopSearch : will search key present in arr and return index
// if key not found return -1
// uses loop
func LoopSearch(arr []int, key int) int {
	var min int
	max := len(arr) - 1
	var mid int
	for max >= min {
		mid = (max + min) / 2
		if arr[mid] == key {
			return mid
		} else if arr[mid] < key {
			min = mid + 1
		} else {
			max = mid - 1
		}
	}
	return -1
}

// RecursionSearch : will search key present in arr and return its index
// if key not found return -1
// uses Recursion
func RecursionSearch(arr []int, key int) int {
	return recSearch(arr, key, 0, len(arr)-1)
}

func recSearch(arr []int, key int, min, max int) int {
	if min > max {
		return -1
	}

	mid := (max + min) / 2
	if arr[mid] == key {
		return mid
	} else if arr[mid] < key {
		min = mid + 1
	} else {
		max = mid - 1
	}
	return recSearch(arr, key, min, max)
}
