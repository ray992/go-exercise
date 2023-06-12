package main

/** 字符串中的变位词 **/

func checkInclusion(s1 string, s2 string) bool {
	var p = [26]int{}
	for i := 0; i < len(s1); i++ {
		p[s1[i]-'a']++
	}
	left, right := 0, len(s1)-1
	for right < len(s2) {
		var p1 = [26]int{}
		startIndex := left
		for ; startIndex <= right; startIndex++ {
			p1[s2[startIndex]-'a']++
		}
		if p1 == p {
			return true
		}
		left++
		right++
	}
	return false
}

func checkInclusion2(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	var p1 = [26]int{}
	var p2 = [26]int{}
	for i := 0; i < len(s1); i++ {
		p1[s1[i]-'a']++
		p2[s2[i]-'a']++
	}
	if p1 == p2 {
		return true
	}
	// 子串比较
	for i := len(s1); i < len(s2); i++ {
		p2[s2[i]-'a']++
		p2[s2[i-len(s1)]-'a']--
		if p1 == p2 {
			return true
		}
	}
	return false
}

func main() {

}
