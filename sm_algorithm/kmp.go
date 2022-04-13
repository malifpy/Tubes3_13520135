package smalgorithm

import (
    "errors"
)

func KMP(text string, pattern string) (int, error) {
    if len(pattern) == 0 {
        return -1, errors.New("Pattern Length is 0")
    } else {
        return kmpAlgo(text, pattern), nil
    }
}

func kmpAlgo(text string, pattern string) int {
	borderArray := kmpBorderArray(pattern)

	textLen := len(text)
	ptrnLen := len(pattern)

	i, j := 0, 0

	for i < textLen {
		if text[i] == pattern[j] {
			if j == ptrnLen-1 {
				return i - (ptrnLen - 1)
			}
			i++
			j++
		} else if j > 0 {
			j = borderArray[j-1]
		} else {
			i++
		}
	}
	return -1
}

func kmpBorderArray(input string) []int {
	strLen := len(input)
	// Array dinamis
	borderArray := make([]int, strLen)
	borderArray[0] = 0

	i, j := 1, 0

	for i < strLen {
		if borderArray[j] == borderArray[i] {
			borderArray[i] = j + 1
			i++
			j++
		} else if j > 0 {
			j = borderArray[j-1]
		} else {
			borderArray[i] = 0
			i++
		}
	}

	return borderArray[:]
}
