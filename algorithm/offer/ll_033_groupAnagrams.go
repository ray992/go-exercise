package main

import "sort"

/** 变位词组 **/

func groupAnagrams(strs []string) [][]string {
	mp := make(map[string][]string)
	for _, v := range strs {
		bytes := []byte(v) // 转成字符数组
		sort.Slice(bytes, func(i, j int) bool {
			return bytes[i] < bytes[j]
		}) // 按照字符顺序排序
		sortedString := string(bytes)
		mp[sortedString] = append(mp[sortedString], v)
	}
	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}