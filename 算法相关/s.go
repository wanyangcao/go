package main

import "fmt"

type X struct {
	i int
}

func (x *X) Increase(n int) {
	x.i += n
}

func (x *X) Decrease(n int) {
	x.i -= n
}

func (x X) String() string {
	fmt.Println("1")
	return fmt.Sprintf("%v", x.i)
}

func main() {
	var x X
	x.Increase(10)
	x.Decrease(5)
	fmt.Println(x) // 1, 5
}
