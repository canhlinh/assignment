package main

import (
	"fmt"
	"strings"
)

func main() {
	result := wheelOfFortune([]string{"green", "red", "blue", "ball", "is", "are", "house", "the"}, []string{
		"t  ", "b  l", "i ", "bl  "}, "t   ba l is  lue")
	fmt.Println(result)
}

type ScoreIndex map[int]int

func (si ScoreIndex) HighestIndexByScore() int {
	highestScore := 0
	bestIndex := 0

	for index, score := range si {
		if score > highestScore {
			bestIndex = index
		}
	}

	return bestIndex
}

func wheelOfFortune(input []string, sugestions1 []string, sugestions2 string) string {
	result := []string{}
	if len(sugestions2) > 0 {
		sugestions1 = mixSuggestions(sugestions1, sugestions2)
	}

	for _, sugestion1 := range sugestions1 {
		scoreIndex := ScoreIndex{}

		for j, word := range input {
			if len(word) == len(sugestion1) {
				for k := 0; k < len(sugestion1); k++ {
					if sugestion1[k] == ' ' {
						continue
					}

					if sugestion1[k] != word[k] {
						scoreIndex[j] = 0
						break
					}

					if sugestion1[k] == word[k] {
						scoreIndex[j]++
					}
				}
			}
		}

		result = append(result, input[scoreIndex.HighestIndexByScore()])
	}

	return strings.Join(result, " ")
}

func mixSuggestions(sugestions1 []string, sugestions2 string) []string {
	chacIndex := 0
	mixed := []string{}

	for _, suggestion1 := range sugestions1 {
		suggestion := make([]rune, len(suggestion1))
		for j, chac := range suggestion1 {

			if chacIndex >= len(sugestions2)-1 || sugestions2[chacIndex] == ' ' {
				suggestion[j] = chac
			} else {
				suggestion[j] = rune(sugestions2[chacIndex])
			}

			chacIndex++
		}

		mixed = append(mixed, string(suggestion))
		chacIndex++
	}

	return mixed
}
