package smalgorithm_test

import (
    "dna-matcher/sm_algorithm"
    "testing"
)

func TestKMPNormal(t * testing.T){
    testText := "abacaabaccabacabaabb"
    testPtrn := "abacab"
    fIdx := smalgorithm.KMP(testText, testPtrn)
    if fIdx != 10 {
        t.Fatalf(`Pattern %s found at index %d`, testPtrn, fIdx)
    }
}

func TestKMPExact(t * testing.T){
    testText := "abacab"
    testPtrn := "abacab"
    fIdx := smalgorithm.KMP(testText, testPtrn)
    if fIdx != 0 {
        t.Fatalf(`Pattern %s found at index %d`, testPtrn, fIdx)
    }
}

func TestKMPShort(t * testing.T){
    testText := "a"
    testPtrn := "abacab"
    fIdx := smalgorithm.KMP(testText, testPtrn)
    if fIdx != -1 {
        t.Fatalf(`Pattern %s found at index %d`, testPtrn, fIdx)
    }
}

func TestKMPExactFirst(t * testing.T){
    testText := "abacabawobduawvdo"
    testPtrn := "abacab"
    fIdx := smalgorithm.KMP(testText, testPtrn)
    if fIdx != 0 {
        t.Fatalf(`Pattern %s found at index %d`, testPtrn, fIdx)
    }
}

func TestKMPExactMiddle(t * testing.T){
    testText := "aklwdbuawabacabaobwduia"
    testPtrn := "abacab"
    fIdx := smalgorithm.KMP(testText, testPtrn)
    if fIdx != 9 {
        t.Fatalf(`Pattern %s found at index %d`, testPtrn, fIdx)
    }
}

func TestKMPExactLast(t * testing.T){
    testText := "aowdvuoavwdlblabacab"
    testPtrn := "abacab"
    fIdx := smalgorithm.KMP(testText, testPtrn)
    if fIdx != 14 {
        t.Fatalf(`Pattern %s found at index %d`, testPtrn, fIdx)
    }
}
