package smalgorithm

import (
	"errors"
)

// Wrapper untuk memastikan pattern gak kosong
func KMPMod(text string, pattern string) (int, float64, error) {
	if len(pattern) == 0 {
		return -1, 0.0, errors.New("pattern length is 0")
	} else {
		// (* 2 / 10) karena toleransinya >= 80%
		fIdx, diff := kmpModAlgo(text, pattern, len(pattern)*2/10)
		return fIdx, diff, nil
	}

}

/**
 * Generalisasi dari Algoritma KMP dimana
 * Algoritma akan menoleransi kesalahan
 * sebanyak `tolerance`
 */
func kmpModAlgo(text string, pattern string, tolerance int) (int, float64) {
	borderArray := kmpBorderArray(pattern)

	textLen := len(text)
	ptrnLen := len(pattern)

	i, j, k, diffCount, minCount, minI := 0, 0, -1, 0, ptrnLen, -1
	// i          variabel index untuk text
	// j          variabel index untuk pattern
	// k          kemunculan pertama perbedaan
	// diffCount  jumlah perbedaan
	// minCount   jumlah perbedaan paling sedikit
	// minI       index text untuk minCount

	for i < textLen && !(j == ptrnLen-1 && diffCount == 0) {
		for i < textLen && j < ptrnLen && diffCount <= tolerance {
			if text[i] != pattern[j] {
				diffCount++
				if k == -1 {
					k = j
				}
			}
			i++
			j++
		}
		if j == ptrnLen && diffCount < minCount {
			minCount = diffCount
			if diffCount <= tolerance {
				minI = i - j
			}
		}
		if i < textLen {
			// Penambahan berlebihan di akhir loop sebelumnya
			// sehingga dikurangi
			i--
			j--
			if diffCount > 0 {
				// Pergi ke perbedaan pertama dengan cara
				// yang sama seperti jika ada kesalahan
				// di algoritma KMP biasa
				if k > 0 {
					i = i - (j - k)
					j = borderArray[k-1]
				} else if j > 0 {
					j = borderArray[j-1]
				} else {
					i++
				}
				// Reset k dan diffCount
				k = -1
				diffCount = 0
			}
		}
	}

	// for i < textLen {
	// 	if text[i] == pattern[j] {
	// 		if j == ptrnLen-1 {
	// 			if diffCount == 0 {
	// 				return i - (ptrnLen - 1), 1 - float64(diffCount)/float64(ptrnLen)
	// 			} else {
	// 				if diffCount < minCount {
	// 					minCount = diffCount
	// 					if diffCount <= tolerance {
	// 						minI = i - ptrnLen
	// 					}
	// 				}
	// 				i, j, k, diffCount = i+1, 0, -1, 0
	// 			}
	// 		} else {
	// 			i++
	// 			j++
	// 		}
	// 	} else {
	// 		diffCount++
	// 		if j == ptrnLen-1 && diffCount < minCount {
	// 			minCount = diffCount
	// 			if diffCount <= tolerance {
	// 				minI = i - ptrnLen
	// 			}
	// 		}
	// 		if diffCount <= tolerance && j != ptrnLen-1 {
	// 			if k == -1 {
	// 				k = j
	// 			}
	// 			i++
	// 			j++
	// 		} else {
	// 			if k > 0 {
	// 				i = i - (j - k)
	// 				j = borderArray[k-1]
	// 			} else if j > 0 {
	// 				j = borderArray[j-1]
	// 			} else {
	// 				i++
	// 			}
	// 			k = -1
	// 			diffCount = 0
	// 		}
	// 	}
	// }
	return minI, 1 - float64(minCount)/float64(ptrnLen)
}
