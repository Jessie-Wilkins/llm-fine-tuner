package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var fit_score = []int{0, 0, 0, 0, 0}

	var prompt_index1 = []int{rand.Intn(5), rand.Intn(5), rand.Intn(5), rand.Intn(5), rand.Intn(5)}

	var prompt_index2 = []int{rand.Intn(5), rand.Intn(5), rand.Intn(5), rand.Intn(5), rand.Intn(5)}

	var fit_achieved = false

	var num_of_rounds = 0

	for !fit_achieved {
		prompt_index1, prompt_index2, fit_score, fit_achieved = fit("real apple", prompt_index1, prompt_index2, fit_score)
		num_of_rounds++
		// for i := 0; i < 5; i++ {

		// 	fmt.Printf("Intermdiate response:  %v %v\n", prompt_array1[prompt_index1[i]], prompt_array2[prompt_index2[i]])

		// }
	}
	fmt.Printf("Num of Rounds: %v\n", num_of_rounds)
	for i := 0; i < 5; i++ {
		fmt.Printf("Final response: %v %v\n", prompt_array1[prompt_index1[i]], prompt_array2[prompt_index2[i]])
	}
}
