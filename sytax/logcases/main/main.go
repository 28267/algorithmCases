package main

import (
	"log"
	"os"
	// "github.com/28267/algorithmCases/sytax/logcases/util"
)

func main() {

	// log.SetPrefix("前缀")
	// res := log.Prefix()
	// log.Println(res)
	// log.Println("打印日志")
	// log.Fatalln("Fatal日志")
	// log.Panicln("Panic日志")
	// log.SetFlags(log.Lshortfile | log.Lmicroseconds | log.Lmsgprefix | log.Ldate | log.Ltime)
	// log.Println("打印日志")
	// log.Fatalln("Fatal日志")
	// util.OutToFile()
	logFile, err := os.OpenFile("D:/GitHubwarehouse/algorithmCases/sytax/logcases/test.log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Println(err)
	}
	log.SetOutput(logFile)
	log.SetPrefix("标准输出到日志文件")
	log.Println("这是一条普通的日志")
}
