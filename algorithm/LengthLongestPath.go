package main

import (
	"fmt"
	"math"
	"strings"
)

func lengthLongestPath(input string) int {
	answer := 0
	lineArray := strings.Split(input, "\n")
	depthMap := make(map[int]int, 8)
	for t := 0; t < len(lineArray); t++ {
		depth := getDepth(lineArray[t])
		if depth == 0 {
			depthMap[depth] = len(lineArray[t])
		} else {
			depthMap[depth] = depthMap[depth-1] + 1 + len(lineArray[t]) - depth
		}
		if strings.Contains(lineArray[t], ".") {
			answer = int(math.Max(float64(answer), float64(depthMap[depth])))
		}
	}
	return answer
}

func getDepth(line string) int {
	depth := 0
	for i := 0; i < len(line); i++ {
		if line[i] == '\t' {
			depth++
		}
	}
	return depth
}

func main() {
	fmt.Println(lengthLongestPath("dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext"))
	fmt.Println(lengthLongestPath("dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext"))
}
