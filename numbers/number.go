package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(123_3245)
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)
	fmt.Println(math.MaxInt32)

	var c1 complex64 = 1 + 2i
	var c2 complex128 = 2 + 3i
	fmt.Println(c1, c2)

	s2 := `avac
sfsaf
csdcas`
	fmt.Println(s2)
	ss := strings.Split("abc|efg|hij", "|")
	fmt.Println(ss)
	fmt.Println(strings.HasPrefix("abc", "ab"))
	fmt.Println(strings.Index("abceesgs", "ee"))
	fmt.Println(strings.Contains("afsafsd", "sad"))
}
