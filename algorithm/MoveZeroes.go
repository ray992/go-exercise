package main

func moveZeroes(nums []int) {
	n := len(nums)
	j := 0
	for i := 0; i < n; i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
	}
	for i := j; i < n; i++ {
		nums[i] = 0
	}
}

func moveZeroes2(nums []int) {
	n := len(nums)
	j := 0
	for i := 0; i < n; i++ {
		if nums[i] != 0 {
			if i > j {
				nums[j] = nums[i]
				nums[i] = 0
			}
			j++
		}
	}
}

func main() {
	moveZeroes([]int{0, 1, 0, 3, 12})
	moveZeroes([]int{0})
	moveZeroes([]int{0, 0, 1}) // 0 0 1 1 0 0
}
