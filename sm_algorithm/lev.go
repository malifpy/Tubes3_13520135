package smalgorithm

func LevDist(strA string, strB string) int {
    return levDistRec(strA, strB, make(map[string]int))
}

func levDistRec(strA string, strB string, cache map[string]int) int {
	concatStr := strA + "|" + strB
	res, found := cache[concatStr]
	if !found {
		if len(strB) == 0 {
			res = len(strA)
		} else if len(strA) == 0 {
			res = len(strB)
		} else if strA[0] == strB[0] {
			res = levDistRec(strA[1:], strB[1:], cache)
		} else {
			res = 1 + Mins(
				levDistRec(strA[1:], strB, cache),
				levDistRec(strA, strB[1:], cache),
				levDistRec(strA[1:], strB[1:], cache),
			)
		}
		cache[concatStr] = res
	}
	return res

}

func Mins(value int, values ...int) int {
	for _, v := range values {
		if v < value {
			value = v
		}
	}
	return value
}
