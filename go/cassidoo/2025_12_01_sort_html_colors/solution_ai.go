package sortHtmlColors

// SortHtmlColorsAI sorts HTML color hex values in ascending order using quicksort
func SortHtmlColorsAI(colors []string) []string {
	// Create a copy to avoid modifying the input
	result := make([]string, len(colors))
	copy(result, colors)

	quicksort(result, 0, len(result)-1)
	return result
}

// quicksort implements the quicksort algorithm
func quicksort(arr []string, low, high int) {
	if low < high {
		// Partition the array and get the pivot index
		pivotIdx := partition(arr, low, high)

		// Recursively sort elements before and after partition
		quicksort(arr, low, pivotIdx-1)
		quicksort(arr, pivotIdx+1, high)
	}
}

// partition takes last element as pivot, places it at correct position,
// and places all smaller elements to left and all greater elements to right
func partition(arr []string, low, high int) int {
	pivot := arr[high]
	i := low - 1 // Index of smaller element

	for j := low; j < high; j++ {
		// If current element is smaller than or equal to pivot
		if arr[j] <= pivot {
			i++
			// Swap arr[i] and arr[j]
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	// Swap arr[i+1] and arr[high] (pivot)
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
