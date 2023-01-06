package order

func Quick[T Ordered](data []T) {
	Quicksort[T](data, 0, len(data)-1)
}
func Quicksort[T Ordered](data []T, low, high int) {
	if low < high {
		pivot := partition(data, low, high)
		Quicksort(data, low, pivot)
		Quicksort(data, pivot+1, high)
	}
}

func partition[T Ordered](data []T, low, high int) int {
	pivot := data[low]
	i := low
	j := high

	for i < j {
		for data[i] <= pivot && i < high {
			i++
		}
		for data[j] > pivot && j > low {
			j--
		}
		if i < j {
			data[i], data[j] = data[j], data[i]
		}
	}

	data[low] = data[j]
	data[j] = pivot
	return j
}
