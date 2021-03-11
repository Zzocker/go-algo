package sorting

// HeapSort : will sort given array using heap sort
func HeapSort(a []int) {
	heapify(a) // nlog(n)
	for i := len(a) - 1; i >= 0; i-- {
		a[0], a[i] = a[i], a[0] // n*log(n)
		moveDown(a, 0, i)
	}
}

func heapify(a []int) {
	for i := len(a) - 1; i >= 0; i-- {
		moveDown(a, i, len(a))
	}
}

func moveDown(a []int, i int, n int) {
	for i < n {
		var tmp int
		if 2*i+2 < n {
			tmp = 2*i + 1
			if a[2*i+2] > a[2*i+1] {
				tmp = 2*i + 2
			}
		} else if 2*i+1 < n {
			tmp = 2*i + 1
		} else {
			break
		}

		if a[tmp] > a[i] {
			a[tmp], a[i] = a[i], a[tmp]
			i = tmp
		} else {
			break
		}
	}
}
