package main

import slice_ "github.com/LXJ0000/basic-go/lib/slice"

func main() {

	a := []int{1, 2, 3, 4, 5}
	b := slice_.Max(a)
	c := slice_.Max(a)

	sum := slice_.Max(a)
	println(b, c, sum)
}
