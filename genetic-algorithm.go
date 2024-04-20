package main

import (
	"math/rand"
	"strings"
)

var prompt_array1 = []string{"fake", "real", "possible", "improbable", "fantastical"}

var prompt_array2 = []string{"banana", "apple", "eggplant", "cranberry", "orange"}

var top_fit_score_values = []int{0, 0}
var top_fit_score_locations = []int{0, 0}

func fit(target string, prompt_index1 []int, prompt_index2 []int, fit_score []int) ([]int, []int, []int, bool) {
	var new_prompt_index1 = []int{0, 0, 0, 0, 0}
	var new_prompt_index2 = []int{0, 0, 0, 0, 0}
	var temp_fit_score = []int{0, 0, 0, 0, 0}
	for prompt_i, _ := range prompt_index1 {

		var actual = prompt_array1[prompt_index1[prompt_i]] + " " + prompt_array2[prompt_index2[prompt_i]]

		if actual != target {
			var sep_target = strings.Split(target, " ")
			var sep_actual = strings.Split(actual, " ")

			for i, s := range sep_actual {
				calculateFitness(sep_target, i, s, temp_fit_score, prompt_i)
				mutate(temp_fit_score, prompt_i, fit_score, new_prompt_index1, new_prompt_index2, prompt_index2, prompt_index1)
			}

		} else {
			return prompt_index1, prompt_index2, fit_score, true
		}
	}
	mate(prompt_index1, prompt_index2)
	return new_prompt_index1, new_prompt_index2, fit_score, false
}

func mate(prompt_index1 []int, prompt_index2 []int) {
	var offspring = []int{0, 0}
	if randomCondition() {
		offspring[0] = prompt_index1[top_fit_score_locations[0]]
		offspring[1] = prompt_index2[top_fit_score_locations[1]]
	} else {
		offspring[0] = prompt_index1[top_fit_score_locations[1]]
		offspring[1] = prompt_index2[top_fit_score_locations[0]]
	}

	var replaceable_prompt = 0

	for replaceable_prompt == prompt_index1[top_fit_score_locations[0]] || replaceable_prompt == prompt_index1[top_fit_score_locations[1]] {
		replaceable_prompt = rand.Intn(5)
	}
	prompt_index1[replaceable_prompt] = offspring[0]
	prompt_index2[replaceable_prompt] = offspring[1]
}

func mutate(temp_fit_score []int, prompt_i int, fit_score []int, new_prompt_index1 []int, new_prompt_index2 []int, prompt_index2 []int, prompt_index1 []int) {
	if temp_fit_score[prompt_i] < fit_score[prompt_i] {
		new_prompt_index1[prompt_i] = rand.Intn(len(prompt_array1))
		new_prompt_index2[prompt_i] = rand.Intn(len(prompt_array2))
	} else {
		fit_score[prompt_i] = temp_fit_score[prompt_i]
		if randomCondition() {
			new_prompt_index1[prompt_i] = rand.Intn(len(prompt_array1))
			new_prompt_index2[prompt_i] = prompt_index2[prompt_i]
		} else {
			new_prompt_index1[prompt_i] = prompt_index1[prompt_i]
			new_prompt_index2[prompt_i] = rand.Intn(len(prompt_array2))
		}
	}
}

func calculateFitness(sep_target []string, i int, s string, temp_fit_score []int, prompt_i int) {
	if sep_target[i] == s {
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
