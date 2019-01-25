package main

import (
	"fmt"
	"os/exec"
	"regexp"
)

func main() {
	fmt.Println(getCpuId())
}

/**
 * 获取电脑CPUId
 */
func getCpuId() string {
	cmd := exec.Command("wmic", "cpu", "get", "ProcessorID")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(string(out))
	str := string(out)
	//匹配一个或多个空白符的正则表达式
	reg := regexp.MustCompile("\\s+")
	str = reg.ReplaceAllString(str, "")
	return str[11:]
}
