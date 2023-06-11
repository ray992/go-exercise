package main

func twoSum(numbers []int, target int) []int {
	length := len(numbers)
	var result []int
	for ii := 0; ii < length; ii++ {
		for jj := ii + 1; jj < length; jj++ {
			if numbers[ii]+numbers[jj] == target {
				result = append(result, ii, jj)
				break
			}
		}
	}
	return result
}

//双指针
func twoSum2(numbers []int, target int) []int {
	length := len(numbers)
	ii, jj := 0, length-1
	var result []int
	for {
		sum := numbers[ii] + numbers[jj]
		if sum == target {
			result = append(result, ii, jj)
			break
		} else if sum < target {
			ii++
		} else {
			jj--
		}
	}
	return result
}

func main() {

}
