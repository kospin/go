package main

import (
	"fmt"
	"math"
	"runtime"
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(合計 int) (x, y int) {
	x = 合計 * 4 / 9
	y = 合計 - x
	return
}

func sqrt(x float64) float64 {
	t, z := 1, float64(1)
	for t <= 10 {
		z = z - (z*z-x)/(2*z)
		t++
	}
	return z
}

var i int
var c, python, java bool

func main() {
	k := 3
	fmt.Println("Hello, world.")
	fmt.Printf("現在你有 %g 個問題\n", math.Nextafter32(2, 3))
	fmt.Println(math.Pi)
	fmt.Println(add(42, 15))
	fmt.Println(swap("你好", "全世界"))
	fmt.Println(split(17))
	fmt.Println(i, c, python, java)
	fmt.Println(k)
	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println(i)
	for i < 10 {
		fmt.Println(i)
		i++
	}
	fmt.Println(sqrt(5))
	switch runtime.GOOS {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.", runtime.GOOS)
	}

}
