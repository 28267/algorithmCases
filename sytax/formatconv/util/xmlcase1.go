package util

import (
	"encoding/xml"
	"fmt"
)

type Student struct {
	Name string
	Age  int
	Job  string
}

func (student *Student) XmlFormatConv() (res []byte) {
	student = &Student{
		Name: "tom",
		Age:  18,
		Job:  "IT",
	}
	res, err := xml.MarshalIndent(student, "", "\t")
	if err != nil {
		fmt.Println("xml.MarshalIndent()-错误", err.Error())
	}
	fmt.Println("序列化后的结果= ", string(res))
	return res
}

func (student *Student) XmlUnFormatConv() {
	student = &Student{
		Name: "",
		Age:  0,
		Job:  "",
	}
	res := student.XmlFormatConv()
	err := xml.Unmarshal(res, student)
	if err != nil {
		fmt.Println("xml.UnMarshal()-错误", err.Error())
	}
	fmt.Println("反序列化后的结果= ", *student)
}
