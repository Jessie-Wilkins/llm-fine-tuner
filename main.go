package main

import "fmt"

func main() {

	var fit_score = 0

	var prompt_index1 = 0

	var prompt_index2 = 0

	for i := 0; i < 10; i++ {
		prompt_index1, prompt_index2, fit_score = fit("real apple", response_array[0], prompt_index1, prompt_index2, fit_score)
		fmt.Printf("Intermdiate response:  %v %v\n", prompt_array1[prompt_index1], prompt_array2[prompt_index2])
	}

	fmt.Printf("Final response: %v %v\n", prompt_array1[prompt_index1], prompt_array2[prompt_index2])
}
