package main

import (
	"fmt"
	"math/rand"
	"strings"
)

var prompt_array1 = []string{
	"Answer the question using the context below.",
	"Let’s think step by step.",
	"Take a deep breath and work through the problem step by step.",
	"Follow my instructions exactly.",
	"Think very carefully.",
	"You are a helpful assistant.",
	`Imagine three different experts are answering this question.
	All experts will write down 1 step of their thinking,
	then share it with the group.
	Then all experts will go on to the next step, etc.
	If any expert realises they're wrong at any point then they leave.`,
	"Explain to me like I’m 11 years old.",
	"Explain to me as if I’m a beginner",
	"Write this using simple English like you are explaining it to a 5-year-old.",
	"I will tip you $500 for a better solution.",
	"You MUST do this.",
	"Do this correctly or else.",
	"You are an expert in this.",
	"Include all the necessary details.",
	"Don't mess up.",
	"This is your task.",
}

var top_fit_score_values = []int{0, 0}
var top_fit_score_locations = []int{0, 0}

var top_responses = []string{"", "", ""}

func fit(target string, prompt string, prompt_index1 []int, prompt_index2 []bool, fit_score []int) ([]int, []bool, []int, bool, int) {
	var new_prompt_index1 = []int{0, 0, 0, 0, 0}
	var is_before_prompt_index = []bool{false, false, false, false, false}
	for prompt_i, _ := range prompt_index1 {

		full_prompt := createFullPrompt(prompt_index2, prompt_i, prompt_index1, prompt)

		var resp1 = promptLLm(full_prompt)

		var resp2 = promptLLm(full_prompt)

		var resp3 = promptLLm(full_prompt)

		// fmt.Printf("Response for prompt1 %v: %v\n\n", prompt_i, resp1.Response)
		// fmt.Printf("Response for prompt2 %v: %v\n\n", prompt_i, resp2.Response)
		// fmt.Printf("Response for prompt3 %v: %v\n\n", prompt_i, resp3.Response)

		// fmt.Printf("Prompt %v is: %v\n\n", prompt_i, full_prompt)

		if (strings.TrimPrefix(resp1.Response, " ") == target &&
			strings.TrimPrefix(resp2.Response, " ") != target &&
			strings.TrimPrefix(resp3.Response, " ") != target) ||
			(strings.TrimPrefix(resp1.Response, " ") != target &&
				strings.TrimPrefix(resp2.Response, " ") == target &&
				strings.TrimPrefix(resp3.Response, " ") != target) ||
			(strings.TrimPrefix(resp1.Response, " ") != target &&
				strings.TrimPrefix(resp2.Response, " ") != target &&
				strings.TrimPrefix(resp3.Response, " ") == target) ||
			(strings.TrimPrefix(resp1.Response, " ") != target &&
				strings.TrimPrefix(resp2.Response, " ") != target &&
				strings.TrimPrefix(resp3.Response, " ") != target) {
			var sep_target = strings.Split(target, " ")
			var sep_actual1 = strings.Split(resp1.Response, " ")
			var sep_actual2 = strings.Split(resp2.Response, " ")
			var sep_actual3 = strings.Split(resp3.Response, " ")

			var temp_fit_score_1 = 0
			var temp_fit_score_2 = 0
			var temp_fit_score_3 = 0
			for i, s := range sep_actual1 {
				temp_fit_score_1 = calculateFitness(sep_target, sep_actual1, temp_fit_score_1, i, s)
			}
			for i, s := range sep_actual2 {
				temp_fit_score_2 = calculateFitness(sep_target, sep_actual2, temp_fit_score_2, i, s)
			}
			for i, s := range sep_actual3 {
				temp_fit_score_3 = calculateFitness(sep_target, sep_actual3, temp_fit_score_3, i, s)
			}

			var average_fit_score = (temp_fit_score_1 + temp_fit_score_2 + temp_fit_score_3) / 3

			assignTopScores(average_fit_score, prompt_i)

			if top_fit_score_values[0] == average_fit_score {
				top_responses[0] = resp1.Response
				top_responses[1] = resp2.Response
				top_responses[2] = resp3.Response
			}

			fit_score[prompt_i] = average_fit_score

			mutate(average_fit_score, prompt_i, fit_score, new_prompt_index1, is_before_prompt_index, prompt_index1, prompt_index2)

		} else {
			return prompt_index1, prompt_index2, fit_score, true, prompt_i
		}
	}
	fmt.Printf("Top Score 1: %v\n\n", top_fit_score_values[0])
	fmt.Printf("Top Score 2: %v\n\n", top_fit_score_values[1])
	fmt.Printf("Top Responses: %v\n\n%v\n\n%v\n\n", top_responses[0], top_responses[1], top_responses[2])
	mate(prompt_index1, prompt_index2)
	return new_prompt_index1, is_before_prompt_index, fit_score, false, -1
}

func assignTopScores(average_fit_score int, prompt_i int) {
	if average_fit_score > top_fit_score_values[0] {
		top_fit_score_values[1] = top_fit_score_values[0]
		top_fit_score_locations[1] = top_fit_score_locations[0]
		top_fit_score_values[0] = average_fit_score
		top_fit_score_locations[0] = prompt_i
	} else if average_fit_score > top_fit_score_values[1] {
		top_fit_score_values[1] = average_fit_score
		top_fit_score_locations[0] = prompt_i
	}
}

func createFullPrompt(prompt_index2 []bool, prompt_i int, prompt_index1 []int, prompt string) string {
	var full_prompt string

	if prompt_index2[prompt_i] {
		full_prompt = prompt_array1[prompt_index1[prompt_i]] + " " + prompt
	} else {
		full_prompt = prompt + " " + prompt_array1[prompt_index1[prompt_i]]
	}
	return full_prompt
}

func mate(prompt_index1 []int, prompt_index2 []bool) {
	var offspring_index = 0
	var offspring_before = false
	if randomCondition() {
		offspring_index = prompt_index1[top_fit_score_locations[0]]
		offspring_before = prompt_index2[top_fit_score_locations[1]]
	} else {
		offspring_index = prompt_index1[top_fit_score_locations[1]]
		offspring_before = prompt_index2[top_fit_score_locations[0]]
	}

	var replaceable_prompt = 0

	for replaceable_prompt == prompt_index1[top_fit_score_locations[0]] || replaceable_prompt == prompt_index1[top_fit_score_locations[1]] {
		replaceable_prompt = rand.Intn(5)
	}
	prompt_index1[replaceable_prompt] = offspring_index
	prompt_index2[replaceable_prompt] = offspring_before
}

func mutate(average_fit_score int, prompt_i int, fit_score []int, new_prompt_index1 []int, new_prompt_index2 []bool, prompt_index1 []int, prompt_index2 []bool) {
	if average_fit_score < fit_score[prompt_i] {
		new_prompt_index1[prompt_i] = rand.Intn(len(prompt_array1))
		new_prompt_index2[prompt_i] = randomCondition()
	} else {
		fit_score[prompt_i] = average_fit_score
		if randomCondition() {
			new_prompt_index1[prompt_i] = rand.Intn(len(prompt_array1))
			new_prompt_index2[prompt_i] = prompt_index2[prompt_i]
		} else {
			new_prompt_index1[prompt_i] = prompt_index1[prompt_i]
			new_prompt_index2[prompt_i] = randomCondition()
		}
	}
}

func calculateFitness(sep_target []string, sep_actual []string, temp_fit_score int, i int, s string) int {
	if i < len(sep_target) && sep_target[i] == s {
		temp_fit_score++
	}

	var joined_string = strings.Join(sep_target, " ")

	if strings.Contains(joined_string, s) {
		temp_fit_score++
	}
	if strings.Contains(strings.ToLower(joined_string), strings.ToLower(s)) {
		temp_fit_score++
	}
	if strings.Contains(strings.TrimSpace(strings.ToLower(joined_string)), strings.TrimSpace(strings.ToLower(s))) {
		temp_fit_score++
	}
	replacements := map[string]string{
		".": "",
		",": "",
		"'": "",
		"?": "",
		"!": "",
		"`": "",
	}
	var noPuncJoinString string
	var noPuncS string
	for old, new := range replacements {
		noPuncJoinString = strings.ReplaceAll(joined_string, old, new)
		noPuncS = strings.ReplaceAll(s, old, new)
	}
	if strings.Contains(noPuncJoinString, noPuncS) {
		temp_fit_score++
	}
	if len(sep_actual) <= len(sep_target) {
		temp_fit_score++
	}

	return temp_fit_score

}

func randomCondition() bool {
	var bool_array = [2]bool{true, false}
	return bool_array[rand.Intn(2)]
}
