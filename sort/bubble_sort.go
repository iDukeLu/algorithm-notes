package sort

type Number interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

// 冒泡排序
func BubbleSort[T Number](arr []T) []T {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j] // 交换元素
			}
		}
	}
	return arr
}

// 冒泡排序分解步骤
func BubbleSortStep[T Number](arr []T) {
	n := len(arr)

	// 第 1 轮，需要比较 n-1 次
	for i := 0; i < n-1; i++ {
		if arr[i] > arr[i+1] {
			arr[i], arr[i+1] = arr[i+1], arr[i]
		}
	}

	// 第 2 轮，需要比较 n-2 次
	for i := 0; i < n-2; i++ {
		if arr[i] > arr[i+1] {
			arr[i], arr[i+1] = arr[i+1], arr[i]
		}
	}

	// 第 3 轮，需要比较 n-3 次
	for i := 0; i < n-3; i++ {
		if arr[i] > arr[i+1] {
			arr[i], arr[i+1] = arr[i+1], arr[i]
		}
	}

	// 第 n-1 轮，需要比较 n-(n-1)=1 次
	for i := 0; i < n-3; i++ {
		if arr[i] > arr[i+1] {
			arr[i], arr[i+1] = arr[i+1], arr[i]
		}
	}

	// 发现规律：第 m 轮，需要比较 n-m 次
	var m int
	for i := 0; i < n-m; i++ {
		if arr[i] > arr[i+1] {
			arr[i], arr[i+1] = arr[i+1], arr[i]
		}
	}

	// 总共需要比较 n-1 轮，且每轮比较的次数减 1，
	// 因为每轮比较后会冒泡最大的数字到最右侧
	for j := 0; j < n-1; j++ { // 每次比较相邻两个数，需要比较 n-1 轮
		for i := 0; i < n-(j+1); i++ { // 每轮比较后，最大的数会冒泡到最右侧，故下一轮比较就会少 1 次
			if arr[i] > arr[i+1] { // 相邻元素比较
				arr[i], arr[i+1] = arr[i+1], arr[i] // 交换元素
			}
		}
	}

}
