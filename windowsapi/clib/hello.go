package main

import "C"

//非常重要，export 表示把go的函数映射到python的函数调用
//如果没有export，那么就不能生成.h文件，python也就无法调用该函数

//export Hello
func Hello() string {
	return "hello world"

}

func main() {}
