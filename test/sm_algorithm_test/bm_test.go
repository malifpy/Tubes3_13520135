package smalgorithm_test

import (
    "dna-matcher/sm_algorithm"
    "testing"
)

func testError(t *testing.T, ptrn string, fIdxExp int, fIdxRes int, err error, errExpNil bool){
    if ((errExpNil && (err != nil)) || (!errExpNil && err == nil)) {
        t.Log(err.Error())
    } else if fIdxRes != fIdxExp {
        t.Fatalf(`Pattern %s found at index %d`, ptrn, fIdxRes)
    }
}

func TestBMNormal(t * testing.T){
    testText := "abacaabaccabacabaabb"
    testPtrn := "abacab"
    fIdx, err := smalgorithm.BM(testText, testPtrn)
    testError(t, testPtrn, 10, fIdx, err, true)
}

func TestBMExact(t * testing.T){
    testText := "abacab"
    testPtrn := "abacab"
    fIdx, err := smalgorithm.BM(testText, testPtrn)
    testError(t, testPtrn, 0, fIdx, err, true)
}

func TestBMShort(t * testing.T){
    testText := "a"
    testPtrn := "abacab"
    fIdx, err := smalgorithm.BM(testText, testPtrn)
    testError(t, testPtrn, -1, fIdx, err, true)
}

func TestBMExactFirst(t * testing.T){
    testText := "abacabawobduawvdo"
    testPtrn := "abacab"
    fIdx, err := smalgorithm.BM(testText, testPtrn)
    testError(t, testPtrn, 0, fIdx, err, true)
}

func TestBMExactMiddle(t * testing.T){
    testText := "aklwdbuawabacabaobwduia"
    testPtrn := "abacab"
    fIdx, err := smalgorithm.BM(testText, testPtrn)
    testError(t, testPtrn, 9, fIdx, err, true)
}

func TestBMExactLast(t * testing.T){
    testText := "aowdvuoavwdlblabacab"
    testPtrn := "abacab"
    fIdx, err := smalgorithm.BM(testText, testPtrn)
    testError(t, testPtrn, 14, fIdx, err, true)
}

func TestBMEmptyText(t * testing.T){
    testText := ""
    testPtrn := "abacab"
    fIdx, err := smalgorithm.BM(testText, testPtrn)
    testError(t, testPtrn, -1, fIdx, err, true)
}

func TestBMEmptyPattern(t * testing.T){
    testText := "muhammad alif putra yasa"
    testPtrn := ""
    fIdx, err := smalgorithm.BM(testText, testPtrn)
    testError(t, testPtrn, -1, fIdx, err, false)
}
