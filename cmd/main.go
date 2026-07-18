package main

import (
	"fmt"
	"math"
)

const (
	PI           = math.Pi
	GREET string = "привет"
)

var (
	a int     = 1
	b bool    = true
	c uint32  = 3
	d int32   = 4
	e rune    = 'f'
	f float32 = 3.14
)

func pickMe(god string, arg2, arg3 int) (int, string) {

	fmt.Println(god, arg2, arg3)
	return arg2 + arg3, god + " пидор"
}
func main() {
	var a = 0
	a, suca := pickMe(GREET, a, int(d))
	fmt.Println(a, suca, float32(0.0001)+float32(0.0001))

}
