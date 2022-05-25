package main

import "fmt"
import "strings"

func main() {
	i := 10
	var (
		string1 string = "hello"
		string2 string = "world"
	)
	fmt.Println(i)
	fmt.Println("s1 + s2 = ", string1+string2)
	var (
		bool1 bool
		a     string
		b     int
		c     uint
		d     float32
		e     float64
	)
	fmt.Println(bool1, a, b, c, d, e)
	f := 1
	g := "a"
	h := true
	i = 20
	fmt.Println(f, g, h)
	pi := &i
	fmt.Println(*pi)
	var string3 string = "o"
	fmt.Println(strings.ContainsAny(string1, string3))
}
