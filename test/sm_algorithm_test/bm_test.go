package smalgorithm_test

import (
	"dna-matcher/sm_algorithm"
	"fmt"
	"testing"
)

// type testCase struct {
//     text    string  // Text, tempat pencarian Pattern
//     pattern string  // Pattern yang akan dicari
//     expIdx  int     // Expected Idx
//     expErr  bool    // Apakah Error nil
// }
//
// var testCases = []testCase {
//     {"abacaabaccabacabaabb",    "abacab", 10, true},
//     {"abacab",                  "abacab", 0,  true},
//     {"a",                       "abacab", -1, true},
//     {"abacabawobduawvdo",       "abacab", 0,  true},
//     {"aklwdbuawabacabaobwduia", "abacab", 9,  true},
//     {"aowdvuoavwdlblabacab",    "abacab", 14, true},
//     {"",                        "abacab", -1, true},
//     {"malifpy",                 "",       -1, false},
// }
//
// func testTC(t *testing.T, resIdx int, resErr error, tc testCase){
//     if (tc.expErr && resErr != nil) || (!tc.expErr && resErr == nil) {
//         t.Error(resErr.Error())
//     } else if resIdx != tc.expIdx && resIdx == -1 {
//         t.Errorf("False Negative")
//     } else if resIdx != tc.expIdx && resIdx != -1 {
//         t.Errorf("False Positive")
//     } else if resIdx == tc.expIdx && resIdx == -1 {
//         t.Log("True  Negative")
//     } else if resIdx == tc.expIdx && resIdx != -1 {
//         t.Log("True  Positive")
//     }
// }
//
// func TestBM(t *testing.T){
//     var resIdx int
//     var resErr error
//
//     for _, el := range testCases {
//         resIdx, resErr = smalgorithm.BM(el.text, el.pattern)
//         testTC(t, resIdx, resErr, el)
//     }
//
// }

func testError(t *testing.T, ptrn string, fIdxExp int, fIdxRes int, err error, errExpNil bool){
    if (errExpNil && (err != nil)){
        t.Error(err.Error())
    } else if (!errExpNil && err == nil)  {
        t.Error("Expected Error, got nil")
    } else if fIdxRes != fIdxExp {
        var strExp string
        var strRes string
        if fIdxExp != -1 {
            strExp = fmt.Sprintf(`Pattern can be found at %d`, fIdxExp)
        } else {
            strExp = "Pattern cannot be found"
        }

        if fIdxRes != -1 {
            strRes = fmt.Sprintf(`Algorithm found at %d`, fIdxRes)
        } else {
            strRes = "Algorithm did not found the pattern"
        }
        t.Errorf(`%s, %s`, strExp, strRes)
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
