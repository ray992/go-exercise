package main

import "fmt"

func predictPartyVictory(senate string) string {
	n := len(senate)
	rn := 0 // 天辉
	dn := 0 // 夜魇
	for i := 0; i < n; i++ {
		if senate[i] == 'R' {
			rn++
		} else if senate[i] == 'D' {
			dn++
		}
	}
	visit := make(map[int]bool, n)
	count := make(map[int]int, 2)
	tag := -1 // 初始化
	for rn > 0 && dn > 0 {
		for i := 0; i < len(senate); i++ {
			if visit[i] {
				continue
			}
			curChar := senate[i]
			if curChar == 'R' {
				if tag == -1 {
					tag = 0
					count[0] = 1
				} else if tag == 0 {
					count[0] = count[0] + 1
				} else if tag == 1 {
					visit[i] = true
					rn--
					count[1] = count[1] - 1
					if count[1] == 0 {
						tag = -1
					}
				}
			} else if curChar == 'D' {
				if tag == -1 {
					tag = 1
					count[1] = 1
				} else if tag == 0 {
					visit[i] = true
					dn--
					count[0] = count[0] - 1
					if count[0] == 0 {
						tag = -1
					}
				} else if tag == 1 {
					count[1] = count[1] + 1
				}
			}
		}
	}
	if rn == 0 {
		return "Dire"
	}
	return "Radiant"
}

func predictPartyVictory2(senate string) string {
	var radiant, dire []int
	for i, s := range senate {
		if s == 'R' {
			radiant = append(radiant, i)
		} else {
			dire = append(dire, i)
		}
	}
	for len(radiant) > 0 && len(dire) > 0 {
		if radiant[0] < dire[0] {
			radiant = append(radiant, radiant[0]+len(senate))
		} else {
			dire = append(dire, dire[0]+len(senate))
		}
		radiant = radiant[1:]
		dire = dire[1:]
	}
	if len(radiant) > 0 {
		return "Dire"
	}
	return "Radiant"
}

func main() {
	//fmt.Println(predictPartyVictory("RD"))
	//fmt.Println(predictPartyVictory("RDD"))
	fmt.Println(predictPartyVictory("DDRRR"))
}
