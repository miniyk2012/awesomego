package main

import "fmt"

func GenSlice() []string {
	z := []string{"a", "b"}
	fmt.Printf("%p %p\n", &z, &z[0])
	return z
}

func Demo() (b []string) {
	b = GenSlice()
	fmt.Printf("%p %p\n", &b, &b[0])
	return
}

func main() {
	c := Demo()
	fmt.Printf("%v \n%p %p\n", c, &c, &c[0])
}
