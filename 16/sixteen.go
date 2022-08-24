package sixteen

import "fmt"

func QuickSort(arr []int) (array []int) {
	if len(arr) < 2 {
		return arr
	}
	var left, right []int

	pivot := arr[len(arr)-1]
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > pivot {
			right = append(right, arr[i])
		}
		if arr[i] <= pivot {
			left = append(left, arr[i])
		}
	}
	array = append(array, QuickSort(left)...)
	array = append(array, pivot)
	array = append(array, QuickSort(right)...)
	return array
}

func main() {
	arr := []int{2, 4, 5, 1, 0, -3, -5, 7, 8, 1, 22, 15, 14, 99, 11}
	fmt.Println(QuickSort(arr))
}
