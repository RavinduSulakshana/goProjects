package main

import "fmt"

func bucketSort(arr []int)[]int{
	maxVal:= 0
	for _,val := range arr {
		if val>maxVal {
			maxVal=val
		}
	}

	buckets := make([][]int,maxVal+1)

	for i := range buckets {
		buckets[i]=make([]int,0)
	}
	
	for _,val := range arr {
		buckets[val]=append(buckets[val],val)
	}

	result := make([]int,0)

	for _, bucket:= range buckets {
		result = append(result,bucket...)
	}

	return result

}
func main() {
arr := []int{67, 32, 12, 54, 43, 57}
fmt.Println("The given unsorted array is:", arr)
sortedArr := bucketSort(arr)
fmt.Println("The obtained sorted array is:", sortedArr)
}