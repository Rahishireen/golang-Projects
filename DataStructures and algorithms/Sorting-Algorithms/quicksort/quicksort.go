package quicksort

func QuickSort(array []int, start int, end int) {
	if start < end {
		pindex := partition(array, start, end)
		QuickSort(array, start, pindex-1)
		QuickSort(array, pindex+1, end)
	}

}

func partition(array []int, start int, end int) int {
	pivot := array[end]
	pindex := start
	var i int
	for i = start; i < end; i++ {
		if array[i] <= pivot {
			array[i],array[pindex] = array[pindex], array[i]
			pindex++
		}
	}

	array[end], array[pindex] = array[pindex], array[end]
	return pindex

}