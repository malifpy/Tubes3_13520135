package smalgorithm_test

import (
	smalgorithm "dna-matcher/sm_algorithm"
	"testing"
)

func TestIsValidNormal(t *testing.T) {
	testString := "ACGTTTGATGCA"
	res, err := smalgorithm.IsValid(testString)
	// Kalau hasilnya false atau error
	if !res || err != nil {
		t.Fatalf(
			`TEST %s | match_result:%t err:%v`,
			testString,
			res,
			err,
		)
	}
}

func TestIsValidEnd(t *testing.T) {
	testString := "ACGTTTGATGCAc"
	res, err := smalgorithm.IsValid(testString)
	// Kalau hasilnya true atau error
	if res || err != nil {
		t.Fatalf(
			`TEST %s | match_result:%t err:%v`,
			testString,
			res,
			err,
		)
	}
}

func TestIsValidMiddle(t *testing.T) {
	testString := "ACGTTbTGATGCA"
	res, err := smalgorithm.IsValid(testString)
	// Kalau hasilnya true atau error
	if res || err != nil {
		t.Fatalf(
			`TEST %s | match_result:%t err:%v`,
			testString,
			res,
			err,
		)
	}
}

func TestIsValidFirst(t *testing.T) {
	testString := "zAAAAAAAAAAAA"
	res, err := smalgorithm.IsValid(testString)
	// Kalau hasilnya true atau error
	if res || err != nil {
		t.Fatalf(
			`TEST %s | match_result:%t err:%v`,
			testString,
			res,
			err,
		)
	}
}

func TestIsValidRandom(t *testing.T) {
	testString := "awidboqb2 eiq0 f ll"
	res, err := smalgorithm.IsValid(testString)
	// Kalau hasilnya true atau error
	if res || err != nil {
		t.Fatalf(
			`TEST %s | match_result:%t err:%v`,
			testString,
			res,
			err,
		)
	}
}

func TestIsValidEmpty(t *testing.T) {
	testString := ""
	res, err := smalgorithm.IsValid(testString)
	// Kalau hasilnya true atau error
	if res || err != nil {
		t.Fatalf(
			`TEST %s | match_result:%t err:%v`,
			testString,
			res,
			err,
		)
	}
}

func TestIsDateNormal(t *testing.T) {
	testString := "2003-06-26"
	res, err := smalgorithm.IsDate(testString)
	// Kalau hasilnya false atau error
	if !res || err != nil {
		t.Fatalf(
			`TEST %s | match_result:%t err:%v`,
			testString,
			res,
			err,
		)
	}
}

func TestIsDateReversed(t *testing.T) {
	testString := "06-26-2003"
	res, err := smalgorithm.IsDate(testString)
	// Kalau hasilnya true atau error
	if res || err != nil {
		t.Fatalf(
			`TEST %s | match_result:%t err:%v`,
			testString,
			res,
			err,
		)
	}
}

func TestIsDateEmpty(t *testing.T) {
	testString := ""
	res, err := smalgorithm.IsDate(testString)
	// Kalau hasilnya true atau error
	if res || err != nil {
		t.Fatalf(
			`TEST %s | match_result:%t err:%v`,
			testString,
			res,
			err,
		)
	}
}

func TestIsDateCharacter(t *testing.T) {
	testString := "YYYY-MM-DD"
	res, err := smalgorithm.IsDate(testString)
	// Kalau hasilnya true atau error
	if res || err != nil {
		t.Fatalf(
			`TEST %s | match_result:%t err:%v`,
			testString,
			res,
			err,
		)
	}
}
