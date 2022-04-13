package smalgorithm_test

import (
    "dna-matcher/sm_algorithm"
    "testing"
)

func TestBMNormal(t * testing.T){
    testText := "abacaabaccabacabaabb"
    testPtrn := "abacab"
    fIdx := smalgorithm.BM(testText, testPtrn)
    if fIdx != 10 {
        t.Fatalf(`Pattern %s found at index %d`, testPtrn, fIdx)
    }
}

func TestBMExact(t * testing.T){
    testText := "abacab"
    testPtrn := "abacab"
    fIdx := smalgorithm.BM(testText, testPtrn)
    if fIdx != 0 {
        t.Fatalf(`Pattern %s found at index %d`, testPtrn, fIdx)
    }
}

func TestBMShort(t * testing.T){
    testText := "a"
    testPtrn := "abacab"
    fIdx := smalgorithm.BM(testText, testPtrn)
    if fIdx != -1 {
        t.Fatalf(`Pattern %s found at index %d`, testPtrn, fIdx)
    }
}

func TestBMExactFirst(t * testing.T){
    testText := "abacabawobduawvdo"
    testPtrn := "abacab"
    fIdx := smalgorithm.BM(testText, testPtrn)
    if fIdx != 0 {
        t.Fatalf(`Pattern %s found at index %d`, testPtrn, fIdx)
    }
}

func TestBMExactMiddle(t * testing.T){
    testText := "aklwdbuawabacabaobwduia"
    testPtrn := "abacab"
    fIdx := smalgorithm.BM(testText, testPtrn)
    if fIdx != 9 {
        t.Fatalf(`Pattern %s found at index %d`, testPtrn, fIdx)
    }
}

func TestBMExactLast(t * testing.T){
    testText := "aowdvuoavwdlblabacab"
    testPtrn := "abacab"
    fIdx := smalgorithm.BM(testText, testPtrn)
    if fIdx != 14 {
        t.Fatalf(`Pattern %s found at index %d`, testPtrn, fIdx)
    }
}
