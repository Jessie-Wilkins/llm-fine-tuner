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

func fit(target string, prompt string, prompt_index1 []int, prompt_index2 []bool, fit_score []int) ([]int, []bool, []int, bool, int) {
	var new_prompt_index1 = []int{0, 0, 0, 0, 0}
	var is_before_prompt_index = []bool{false, false, false, false, false}
	var temp_fit_score = []int{0, 0, 0, 0, 0}
	for prompt_i, _ := range prompt_index1 {

		full_prompt := createFullPrompt(prompt_index2, prompt_i, prompt_index1, prompt)

		var resp = promptLLm(full_prompt)

		fmt.Printf("Response for prompt %v: %v\n\n", prompt_i, resp.Response)

		fmt.Printf("Prompt %v is: %v\n\n", prompt_i, full_prompt)

		if strings.TrimPrefix(resp.Response, " ") != target {
			var sep_target = strings.Split(target, " ")
			var sep_actual = strings.Split(resp.Response, " ")

			for i, s := range sep_actual {
				calculateFitness(sep_target, i, s, temp_fit_score, prompt_i)
				mutate(temp_fit_score, prompt_i, fit_score, new_prompt_index1, is_before_prompt_index, prompt_index1, prompt_index2)
			}

		} else {
			return prompt_index1, prompt_index2, fit_score, true, prompt_i
		}
	}
	mate(prompt_index1, prompt_index2)
	return new_prompt_index1, is_before_prompt_index, fit_score, false, -1
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

func mutate(temp_fit_score []int, prompt_i int, fit_score []int, new_prompt_index1 []int, new_prompt_index2 []bool, prompt_index1 []int, prompt_index2 []bool) {
	if temp_fit_score[prompt_i] < fit_score[prompt_i] {
		new_prompt_index1[prompt_i] = rand.Intn(len(prompt_array1))
		new_prompt_index2[prompt_i] = randomCondition()
	} else {
		fit_score[prompt_i] = temp_fit_score[prompt_i]
		if randomCondition() {
			new_prompt_index1[prompt_i] = rand.Intn(len(prompt_array1))
			new_prompt_index2[prompt_i] = prompt_index2[prompt_i]
		} else {
			new_prompt_index1[prompt_i] = prompt_index1[prompt_i]
			new_prompt_index2[prompt_i] = randomCondition()
		}
	}
}

func calculateFitness(sep_target []string, i int, s string, temp_fit_score []int, prompt_i int) {
	if i < len(sep_target) && sep_target[i] == s {
		temp_fit_score[prompt_i]++
	}

	var joined_string = strings.Join(sep_target, " ")

	if strings.Contains(joined_string, s) {
		temp_fit_score[prompt_i]++
	}
	if strings.Contains(strings.ToLower(joined_string), strings.ToLower(s)) {
		temp_fit_score[prompt_i]++
	}
	if strings.Contains(strings.TrimSpace(strings.ToLower(joined_string)), strings.TrimSpace(strings.ToLower(s))) {
		temp_fit_score[prompt_i]++
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
		temp_fit_score[prompt_i]++
	}

	if temp_fit_score[prompt_i] > top_fit_score_values[0] {
		top_fit_score_values[1] = top_fit_score_values[0]
		top_fit_score_locations[1] = top_fit_score_locations[0]
		top_fit_score_values[0] = temp_fit_score[prompt_i]
		top_fit_score_locations[0] = prompt_i
	} else if temp_fit_score[prompt_i] > top_fit_score_values[1] {
		top_fit_score_values[1] = temp_fit_score[prompt_i]
		top_fit_score_locations[0] = prompt_i
	}

}

func randomCondition() bool {
	var bool_array = [2]bool{true, false}
	return bool_array[rand.Intn(2)]
}
