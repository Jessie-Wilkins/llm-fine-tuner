package main

import "fmt"

func main() {

	var fit_score = 0

	var prompt_index1 = []int{0, 0, 0, 0, 0}

	var prompt_index2 = []int{0, 0, 0, 0, 0}

	for i := 0; i < 10; i++ {
		prompt_index1, prompt_index2, fit_score = fit("real apple", prompt_index1[:], prompt_index2[:], fit_score[:])
		for i := 0; i < 5; i++ {

			fmt.Printf("Intermdiate response:  %v %v\n", prompt_array1[prompt_index1[i]], prompt_array2[prompt_index2[i]])

		}
	}

	for i := 0; i < 5; i++ {
		fmt.Printf("Final response: %v %v\n", prompt_array1[prompt_index1[i]], prompt_array2[prompt_index2[i]])
	}
}
