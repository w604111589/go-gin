package example

import (
	"fmt"
	"log"
	"os"
)

//LogFatal ...
func LogFatal(s string) {
	log.Fatal("这是一个致命的错误")
}

//LogFatal1 ...
func LogFatal1() {
	file, err := os.OpenFile("test.log", os.O_APPEND|os.O_CREATE, 666)
	if err != nil {
		log.Fatalln("出现了未知的错误", err)
	}
	logger := log.New(file, "wt_", log.LstdFlags|log.Llongfile)
	fmt.Println("将内容写入命令行终端")
	logger.Println("将内容写入日志文件中")

	// logger.Fatalln("输出数据，并中断执行")
	log.Panicln("输出数据，并中断执行")
	// panic("恐慌")
}
