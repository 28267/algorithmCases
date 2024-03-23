package util

import (
	"flag"
	"fmt"
)

// func ParseParams() {
// 	name := flag.String("name", "sakura", "姓名")
// 	job := flag.String("job", "IT", "工作")
// 	sex := flag.Bool("sex", true, "性别")
// 	flag.Parse()
// 	fmt.Println(*name, *job, *sex)
// }

func ParseVarParams() {
	var name string
	var job string
	var sex bool
	flag.StringVar(&name, "name", "sakura", "姓名")
	flag.StringVar(&job, "job", "IT", "工作")
	flag.BoolVar(&sex, "sex", true, "性别")
	flag.Parse()
	fmt.Println(name, job, sex)
}
