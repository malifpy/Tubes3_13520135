package smalgorithm

import (
    "errors"
)

func BM(text string, pattern string) (int, error) {
    if len(pattern) == 0 {
        return -1, errors.New("Pattern Length is 0")
    } else {
        return bmAlgo(text, pattern), nil
    }
}

func bmAlgo(text string, pattern string) int {
    lastOccur := bmLastOccur(pattern)

    textLen := len(text)
    ptrnLen := len(pattern)

    i := ptrnLen - 1

    if i > textLen - 1 {
        return -1
    }

    j := ptrnLen - 1

    // Emulasi do while loop
    var lastOcc int
    var exists bool
    firstItr := true
    for (firstItr || i <= textLen - 1) {
        firstItr = false
        if pattern[j] == text[i]{
            if j == 0 {
                return i
            }
            i--
            j--
        } else {
            lastOcc, exists = lastOccur[text[i]]
            if !exists {
                lastOcc = -1
            }
            i = i + ptrnLen - Mins(j, 1 + lastOcc)
            j = ptrnLen - 1
        }
    }
    return -1
}

func bmLastOccur(pattern string) map[byte]int {
    lastOccur := make(map[byte]int)
    ptrnLen := len(pattern)

    for i := 0; i < ptrnLen; i++ {
        lastOccur[pattern[i]] = i
    }

    return lastOccur
}

// func min(a int, b int) int {
//     if a < b {
//         return a
//     }
//
//     return b
// }
