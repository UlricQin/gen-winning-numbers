package utils

func IntsIntersection(multiInts map[string][]int, needCnt []int) []int {
	needCntLen := len(needCnt)
	needCntLenLess1 := needCntLen - 1

	var base map[int]int = make(map[int]int)
	for i := 0; i < 1000; i++ {
		base[i] = 0
	}

	for _, ints := range multiInts {
		sz := len(ints)
		for i := 0; i < sz; i++ {
			base[ints[i]] += 1
		}
	}

	ret := make([]int, 0)
	for i := 0; i < 1000; i++ {
		occurNum := base[i]
		if occurNum < needCnt[0] || occurNum > needCnt[needCntLenLess1] {
			continue
		}

		for j := 0; j < needCntLen; j++ {
			if occurNum == needCnt[j] {
				ret = append(ret, i)
			}
		}
	}

	return ret
}
