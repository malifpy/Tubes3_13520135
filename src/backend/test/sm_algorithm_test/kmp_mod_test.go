package smalgorithm_test

import (
    "testing"
    "dna-matcher/sm_algorithm"
)
//
// func TestKMPModAlgoNormal(t *testing.T){
//     fmt.Println(smalgorithm.KMPMod("muhammad alifpyutra yasa", "malifpy"))
// }
func TestKMPModNormal(t * testing.T){
    testText := "abacaabaccabacabaabb"
    testPtrn := "abacab"
    fIdx, _, err:= smalgorithm.KMPMod(testText, testPtrn)
    testError(t, testPtrn, 10, fIdx, err, true)
}

func TestKMPModExact(t * testing.T){
    testText := "abacab"
    testPtrn := "abacab"
    fIdx, _, err:= smalgorithm.KMPMod(testText, testPtrn)
    testError(t, testPtrn, 0, fIdx, err, true)
}

func TestKMPModShort(t * testing.T){
    testText := "a"
    testPtrn := "abacab"
    fIdx, _, err:= smalgorithm.KMPMod(testText, testPtrn)
    testError(t, testPtrn, -1, fIdx, err, true)
}

func TestKMPModExactFirst(t * testing.T){
    testText := "abacabawobduawvdo"
    testPtrn := "abacab"
    fIdx, _, err:= smalgorithm.KMPMod(testText, testPtrn)
    testError(t, testPtrn, 0, fIdx, err, true)
}

func TestKMPModExactMiddle(t * testing.T){
    testText := "aklwdbuawabacabaobwduia"
    testPtrn := "abacab"
    fIdx, _, err:= smalgorithm.KMPMod(testText, testPtrn)
    testError(t, testPtrn, 9, fIdx, err, true)
}

func TestKMPModExactLast(t * testing.T){
    // testText := "aowdvuoavwdlblabacab"
    testText := "aabacab"
    testPtrn := "abacab"
    fIdx, _, err:= smalgorithm.KMPMod(testText, testPtrn)
    testError(t, testPtrn, 1, fIdx, err, true)
}

func TestKMPModEmptyText(t * testing.T){
    testText := ""
    testPtrn := "abacab"
    fIdx, _, err:= smalgorithm.KMPMod(testText, testPtrn)
    testError(t, testPtrn, -1, fIdx, err, true)
}

func TestKMPModEmptyPattern(t * testing.T){
    testText := "muhammad alif putra yasa"
    testPtrn := ""
    fIdx, _, err:= smalgorithm.KMPMod(testText, testPtrn)
    testError(t, testPtrn, -1, fIdx, err, false)
}
//
func TestKMPModSimiliarMiddle(t * testing.T){
    testText := "malifpy"
    testPtrn := "alir"
    fIdx, _, err:= smalgorithm.KMPMod(testText, testPtrn)
    testError(t, testPtrn, -1, fIdx, err, true)
}

func TestKMPModSimiliarPerfect(t * testing.T){
    testText := "malyfpymalifpy"
    testPtrn := "malifpy"
    fIdx, _, err:= smalgorithm.KMPMod(testText, testPtrn)
    testError(t, testPtrn, 7, fIdx, err, true)
}
