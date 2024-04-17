package main

import (
	"math/rand"
	"strings"
)

var response_array = [5]string{"fake banana", "real apple", "possible eggplant", "improbable cranberry", "fantastical orange"}

var prompt_array1 = [5]string{"fake", "real", "possible", "improbable", "fantastical"}

var prompt_array2 = [5]string{"banana", "apple", "eggplant", "cranberry", "orange"}

func fit(target string, actual string, prompt_index1 int, prompt_index2 int, fit_score int) (int, int, int) {
	var new_prompt_index1 int = 0
	var new_prompt_index2 int = 0
	if actual != target {
		var sep_target = strings.Split(target, " ")
		var sep_actual = strings.Split(actual, " ")
		var temp_fit_score = 0
		for i, s := range sep_actual {
			if sep_target[i] == s {
				temp_fit_score++
			}
		}
		if temp_fit_score < fit_score {
			new_prompt_index1 = rand.Intn(len(prompt_array1))
			new_prompt_index2 = rand.Intn(len(prompt_array2))
		} else {
			var bool_array = [2]bool{true, false}
			fit_score = temp_fit_score
			if bool_array[rand.Intn(1)] {
				new_prompt_index1 = rand.Intn(len(prompt_array1))
				new_prompt_index2 = prompt_index2
			} else {
				new_prompt_index1 = prompt_index1
				new_prompt_index2 = rand.Intn(len(prompt_array2))
			}
		}
		return new_prompt_index1, new_prompt_index2, fit_score

	} else {
		return prompt_index1, prompt_index2, fit_score
	}
}
