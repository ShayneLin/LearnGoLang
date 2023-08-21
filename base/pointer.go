package main

import "fmt"

func swap(a *int, b *int) {
	var temp int
	temp = *a
	*a = *b
	*b = temp
}

func main() {
	var a = 10
	var b = 20
	swap(&a, &b)
	fmt.Println("a=", a)
	fmt.Println("b=", b)

	var p *int
	p = &a
	fmt.Println(&a)
	fmt.Println(p)
	fmt.Println(*p)

	var pp **int
	pp = &p
	fmt.Println(&p)
	fmt.Println(pp)
	fmt.Println(**pp)
}
