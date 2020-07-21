package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
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
	z := float64(1)
	for t := 1; t <= 10; t++ {
		z = z - (平方(z)-x)/(2*z)
	}
	return z
}

func 平方(x float64) float64 {
	return x * x
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
		fmt.Printf("%s.\n", runtime.GOOS)
	}

	defer fmt.Println("自爆...") //延遲啊...好像很有用
	defer fmt.Println("我先!!")  //後進先出

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("早安")
	default:
		fmt.Println("晚安")
	}
	var p *int //指標!!
	n := 42
	p = &n
	fmt.Println(*p) //沒有指標運算,那有什麼用?

	type 向量 struct {
		X int
		Y int
	}

	fmt.Println(向量{1, 2})
	v := 向量{3, 4}
	w := &v
	w.X = 1e9 //w.X 等同v.X
	fmt.Println(v, w)

	a := [2]string{"hello", "world"}
	fmt.Println(a)

	primes := []int{2, 3, 5, 7, 11, 13}
	s := primes[1:4]
	s[0] = 23
	fmt.Println(s, primes) //切片不是新的矩陣

}
