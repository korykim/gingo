package main

import "fmt"

func main() {
	a := 20

	var ip *int

	ip = &a

	fmt.Printf("value a address: %x\n", &a)

	fmt.Printf("value ip address: %x\n", ip)

	fmt.Printf("*ip 变量的值: %d\n", *ip)

	var ptr *int

	if ptr != nil {
		fmt.Print("ptr 不是空指针\n")
	} else {
		fmt.Print("ptr 是空指针\n")
	}

}
