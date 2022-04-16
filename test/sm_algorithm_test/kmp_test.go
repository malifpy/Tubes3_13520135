package smalgorithm_test

import (
    "dna-matcher/sm_algorithm"
    "testing"
)

func TestKMPNormal(t * testing.T){
    testText := "abacaabaccabacabaabb"
    testPtrn := "abacab"
    fIdx, err := smalgorithm.KMP(testText, testPtrn)
    testError(t, testPtrn, 10, fIdx, err, true)
}

func TestKMPExact(t * testing.T){
    testText := "abacab"
    testPtrn := "abacab"
    fIdx, err := smalgorithm.KMP(testText, testPtrn)
    testError(t, testPtrn, 0, fIdx, err, true)
}

func TestKMPShort(t * testing.T){
    testText := "a"
    testPtrn := "abacab"
    fIdx, err := smalgorithm.KMP(testText, testPtrn)
    testError(t, testPtrn, -1, fIdx, err, true)
}

func TestKMPExactFirst(t * testing.T){
    testText := "abacabawobduawvdo"
    testPtrn := "abacab"
    fIdx, err := smalgorithm.KMP(testText, testPtrn)
    testError(t, testPtrn, 0, fIdx, err, true)
}

func TestKMPExactMiddle(t * testing.T){
    testText := "aklwdbuawabacabaobwduia"
    testPtrn := "abacab"
    fIdx, err := smalgorithm.KMP(testText, testPtrn)
    testError(t, testPtrn, 9, fIdx, err, true)
}

func TestKMPExactLast(t * testing.T){
    testText := "aowdvuoavwdlblabacab"
    testPtrn := "abacab"
    fIdx, err := smalgorithm.KMP(testText, testPtrn)
    testError(t, testPtrn, 14, fIdx, err, true)
}

func TestKMPEmptyText(t * testing.T){
    testText := ""
    testPtrn := "abacab"
    fIdx, err := smalgorithm.KMP(testText, testPtrn)
    testError(t, testPtrn, -1, fIdx, err, true)
}

func TestKMPEmptyPattern(t * testing.T){
    testText := "muhammad alif putra yasa"
    testPtrn := ""
    fIdx, err := smalgorithm.KMP(testText, testPtrn)
    testError(t, testPtrn, -1, fIdx, err, false)
}
