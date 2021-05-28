package mergesort

//MergeSort
func Mergesort(originalArray []int) []int {
	arrayLen:= len(originalArray)
	 if arrayLen == 1 {
		  return originalArray
	 } else {
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
	return Merge(Mergesort(leftArray),Mergesort(rightArray))
}

}

func Merge(left []int, right []int) []int{	
	leftArrayLen:=len(left)
	rightArrayLen:=len(right)
	var result []int
	i:=0
	j:=0
	k:=0
	for (i<leftArrayLen && j<rightArrayLen){
		if (left[i]<right[j]){
			result=append(result,left[i])
			
			i++			
		}else{
			result=append(result,right[j])
			j++		
		}
		k++
	}

	for (i<leftArrayLen){
		result=append(result,left[i])
		i++
		k++
	}
	for (j<rightArrayLen){
		result=append(result,right[j])
		j++
		k++
	}
	return result
}