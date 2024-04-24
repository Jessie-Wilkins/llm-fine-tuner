package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
)

type PromptResult struct {
	Prompt string `json:"prompt"`
	Result string `json:"result"`
}

func main() {

	var content, err = os.ReadFile("prompt-result.json")
	if err != nil {
		log.Fatal(err)
	}

	var prompt_result_obj PromptResult

	err = json.Unmarshal(content, &prompt_result_obj)

	if err != nil {
		fmt.Println("JSON decode error!")
		return
	}

	var fit_score = []int{0, 0, 0, 0, 0}

	var prompt_index1 = []int{rand.Intn(5), rand.Intn(5), rand.Intn(5), rand.Intn(5), rand.Intn(5)}

	var prompt_index2 = []bool{false, false, false, false, false}

	var fit_achieved = false

	var num_of_rounds = 0

	var fittest_prompt_index = -1

	for !fit_achieved {

		prompt_index1, prompt_index2, fit_score, fit_achieved, fittest_prompt_index = fit(
			prompt_result_obj.Result,
			prompt_result_obj.Prompt,
			prompt_index1,
			prompt_index2,
			fit_score)

		num_of_rounds++
		// for i := 0; i < 5; i++ {

		// 	fmt.Printf("Intermdiate response:  %v %v\n", prompt_array1[prompt_index1[i]], prompt_array2[prompt_index2[i]])

		// }
	}
	fmt.Printf("Num of Rounds: %v\n", num_of_rounds)

	fmt.Printf("Final response: %v\n", createFullPrompt(prompt_index2, fittest_prompt_index, prompt_index1, prompt_result_obj.Prompt))

}
