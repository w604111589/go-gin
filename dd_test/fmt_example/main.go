package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	StringPlus()
	StringJoin()
	StringBuilder()
	StringBuffer()
	StringFmt()
}

//StringPlus ...
func StringPlus() {
	//字符串拼接 (2)
	str := "我" + "是"
	fmt.Println(str)
}

//StringJoin ...
func StringJoin() {
	strSlice := []string{"我", "是"}
	str := strings.Join(strSlice, "")
	fmt.Println(str)
}

//StringBuilder ...
func StringBuilder() {
	var strBuilder strings.Builder
	strBuilder.WriteString("我")
	strBuilder.WriteString("是")
	fmt.Println(strBuilder.String())
}

//StringBuffer ...
func StringBuffer() {
	var strBuffer bytes.Buffer
	strBuffer.WriteString("我")
	strBuffer.WriteString("是")
	fmt.Println(strBuffer.String())
}

//StringFmt ...
func StringFmt() {
	var str = fmt.Sprintf("%s%s", "我", "是")
	fmt.Println(str)
}
