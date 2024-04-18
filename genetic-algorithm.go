package main

import (
	"math/rand"
	"strings"
)

var prompt_array1 = []string{"fake", "real", "possible", "improbable", "fantastical"}

var prompt_array2 = []string{"banana", "apple", "eggplant", "cranberry", "orange"}

func fit(target string, prompt_index1 []int, prompt_index2 []int, fit_score []int) ([]int, []int, []int) {
	var new_prompt_index1 = []int{0, 0, 0, 0, 0}
	var new_prompt_index2 = []int{0, 0, 0, 0, 0}
	var temp_fit_score = []int{0, 0, 0, 0, 0}
	for prompt_i, _ := range prompt_index1 {

		var actual = prompt_array1[prompt_index1[prompt_i]] + " " + prompt_array2[prompt_index2[prompt_i]]

		if actual != target {
			var sep_target = strings.Split(target, " ")
			var sep_actual = strings.Split(actual, " ")

			for i, s := range sep_actual {
				if sep_target[i] == s {
					temp_fit_score[i]++
				}
				if temp_fit_score[prompt_i] < fit_score[prompt_i] {
					new_prompt_index1[prompt_i] = rand.Intn(len(prompt_array1))
					new_prompt_index2[prompt_i] = rand.Intn(len(prompt_array2))
				} else {
					var bool_array = [2]bool{true, false}
					fit_score[prompt_i] = temp_fit_score[prompt_i]
					if bool_array[rand.Intn(2)] {
						new_prompt_index1[prompt_i] = rand.Intn(len(prompt_array1))
						new_prompt_index2[prompt_i] = prompt_index2[prompt_i]
					} else {
						new_prompt_index1[prompt_i] = prompt_index1[prompt_i]
						new_prompt_index2[prompt_i] = rand.Intn(len(prompt_array2))
					}
				}
			}

		} else {
			return prompt_index1, prompt_index2, fit_score
		}
	}
	return new_prompt_index1, new_prompt_index2, fit_score
}
