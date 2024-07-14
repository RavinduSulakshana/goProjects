package main

import "fmt"

func findLargest(numbers []int) int{
	if len(numbers) == 0 {
		return 0
	}
	largest := numbers[0]
	for _, number := range numbers {
		if(number>largest){
			largest=number
		}
}
	return largest
}

func main(){
	numbers := []int{1,2,65,4,99,88,54,0}
	largest := findLargest(numbers)
	fmt.Println("largest Number is = ",largest)


}