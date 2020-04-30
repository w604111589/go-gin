package main

import(
	"strings"
	"fmt"
)

var (
	dingTalkURL = "https://oapi.dingtalk.com/robot/send"
	accessToken = "e9b1ea2a4e71fe2ee21e9dba12917582c6234c520dc410c12f0e9ddd7fb689ef"
)

func main(){
	//字符串包含
	status := strings.Contains("ab","bc") 
	fmt.Println(status)

	num := strings.Count("tttabca","a")
	fmt.Println(num)

	status1 := strings.HasPrefix("https://www.baidu.com","https://")
	fmt.Println(status1)

	status2 := strings.TrimPrefix("https://www.baidu.com","https://")
	fmt.Println(status2)
}