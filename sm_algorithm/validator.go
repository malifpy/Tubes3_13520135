package smalgorithm

import (
	"regexp"
)

func IsValid(s string) (bool, error) {
	var dnaSanitized bool
	var err error

	// Gak menerima string kosong
	if s == "" {
		return false, nil
	}

	// Apakah ada yang bukan ACGT
	dnaSanitized, err = regexp.MatchString("[^ACGT]", s)
	return (!dnaSanitized), err
}
