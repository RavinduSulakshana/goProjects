package main

import "fmt"

func linearSearch (numbers []int,target int) int {
	for index,number := range numbers{
		if number == target{
			return index
		}
	}
	return -1 //target is not found
}

func main(){
	//list of numbers
	numbers := []int{34,46,23,6,2,453,4,64,63}

	//taget number
	target := 453

	//call the function
	result := linearSearch(numbers,target)

	if result != -1 {
        fmt.Printf("Target %d found at index %d.\n", target, result)
    } else {
        fmt.Printf("Target %d not found in the list.\n", target)
    }
}