package main

import "fmt"

func asteroidCollision(asteroids []int) []int {
	n := len(asteroids)
	var res []int
	for i := 0; i < n; i++ {
		aster := asteroids[i]
		if len(res) == 0 {
			res = append(res, aster)
		} else {
			for len(res) > 0 {
				last := res[len(res)-1]
				if last < 0 {
					res = append(res, aster)
					break
				} else if last > 0 {
					if aster > 0 {
						res = append(res, aster)
						break
					} else {
						if last > -aster {
							break
						} else if last == -aster {
							res = res[0 : len(res)-1]
							break
						} else {
							res = res[0 : len(res)-1]
							if len(res) == 0 {
								res = append(res, aster)
								break
							}
						}
					}
				}
			}
		}
	}
	return res
}

func main() {
	fmt.Println(asteroidCollision([]int{5, 10, -5}))
	fmt.Println(asteroidCollision([]int{8, -8}))
	fmt.Println(asteroidCollision([]int{10, 2, -5}))
	fmt.Println(asteroidCollision([]int{-2, -1, 1, 2}))
	fmt.Println(asteroidCollision([]int{1, -2, -2, -2}))
}