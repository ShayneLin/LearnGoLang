package main

import "fmt"

// const 来定义枚举类型
const (
	//可以在const()添加一个关键字iota,每行的iota都会累加1，第一行的iota的默认值是0
	BEIJING  = iota //iota = 0
	SHANGHAI        //iota=1
	SHENZHEN        //iota=2
)

func main() {
	fmt.Printf("BEJING=%d", BEIJING)
	fmt.Printf("SHANGHAI=%d", SHANGHAI)
	fmt.Printf("SHENZHEN=%d", SHENZHEN)
}
