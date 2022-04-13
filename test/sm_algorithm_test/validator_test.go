package smalgorithm_test

import (
    "dna-matcher/sm_algorithm"
    "testing"
)

func TestIsValidNormal(t *testing.T){
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

func TestIsValidEnd(t *testing.T){
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

func TestIsValidMiddle(t *testing.T){
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

func TestIsValidFirst(t *testing.T){
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

func TestIsValidRandom(t *testing.T){
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

func TestIsValidEmpty(t *testing.T){
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
