package main

import "fmt"

func isAlienSorted(words []string, order string) bool {
	orderMap := make(map[byte]int)
	for ii := 0; ii < len(order); ii++ {
		orderMap[order[ii]] = ii
	}
	for ii := 0; ii < len(words); ii++ {
		curWord := words[ii]
		for jj := ii + 1; jj < len(words); jj++ {
			nextWord := words[jj]
			l1, l2 := len(curWord), len(nextWord)
			ll := 0
			if l1 < l2 {
				ll = l1
			} else {
				ll = l2
			}
			normalStatus := false
			for kk := 0; kk < ll; kk++ {
				v1 := orderMap[curWord[kk]]
				v2 := orderMap[nextWord[kk]]
				if v1 < v2 {
					normalStatus = true
					break
				} else if v1 > v2 {
					return false
				}
			}
			if !normalStatus && l2 < l1 {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(isAlienSorted([]string{"word", "world", "row"}, "worldabcefghijkmnpqstuvxyz"))
	fmt.Println(isAlienSorted([]string{"apple", "app"}, "abcdefghijklmnopqrstuvwxyz"))
	fmt.Println(isAlienSorted([]string{"hello", "leetcode"}, "hlabcdefgijkmnopqrstuvwxyz"))
	fmt.Println(isAlienSorted([]string{"fxasxpc", "dfbdrifhp", "nwzgs", "cmwqriv", "ebulyfyve", "miracx", "sxckdwzv", "dtijzluhts", "wwbmnge", "qmjwymmyox"}, "zkgwaverfimqxbnctdplsjyohu"))
}
