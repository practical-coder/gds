package order

type Ordered interface {
	~string | ~int | ~float64
}

func Bubble[T Ordered](data []T) {
	n := len(data)
	for i := 0; i < n - 1; i++ {
		for j := i + 1; j < n; j++ {
			if data[i] > data[j] {
				data[i], data[j] = data[j], data[i]
			}
		}
	}
}

func BubbleBook[T Ordered](data []T) {
	n := len(data)
	for i := 0; i < n - 1; i++ {
		for j := 0; j < n - 1 - i; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
}
