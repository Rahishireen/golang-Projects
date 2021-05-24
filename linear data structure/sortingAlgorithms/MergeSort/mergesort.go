package mergesort

//MergeSort
func Mergesort(originalArray []int) []int {
	arrayLen:= len(originalArray)
	 if arrayLen == 1 {
		 return originalArray
	 }

	mid:= arrayLen/2
	
	leftArray := make([]int,mid)
	rightArray :=  make([]int,(arrayLen-mid))

	for i:=0;i<arrayLen;i++{
	if i<mid{
		leftArray[i]=originalArray[i]
	} else{
		rightArray[i-mid]=originalArray[i]
	}
	}
	leftArray = originalArray[:mid]
	rightArray =originalArray[mid:]
	return merge(leftArray,rightArray,originalArray)

}

func merge(left []int, right []int, original []int) []int{	
	leftArrayLen:=len(left)
	rightArrayLen:=len(right)

	i:=0
	j:=0
	k:=0

	for (i<leftArrayLen && j<rightArrayLen){
		if (left[i]<right[j]){
			original[k]=left[i]
			i++			
		}else{
			original[k]=right[j]
			j++		
		}
		k++
	}

	for (i<leftArrayLen){
		original[k]=left[i]
		i++
		k++
	}
	for (j<rightArrayLen){
		original[k]=right[j]
		j++
		k++
	}
	return original
}