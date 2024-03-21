package util

import (
	"fmt"
)

func MiGong(MiGongArr *[10][10]int, m int, n int) bool {
	//0表示可以探测	1表示是墙 2表示是通路 3表示探测过

	if MiGongArr[7][8] == 2 {
		return true
	} else {
		if MiGongArr[m][n] == 0 {
			MiGongArr[m][n] = 2
			if MiGong(MiGongArr, m+1, n) { //下
				return true
			} else if MiGong(MiGongArr, m, n+1) { //右
				return true
			} else if MiGong(MiGongArr, m, n-1) { //上
				return true
			} else if MiGong(MiGongArr, m-1, n) { //左
				return true
			} else {
				MiGongArr[m][n] = 3
				return false
			}
		} else {
			// MiGongArr[m][n] = 1
			return false //跳出的是递归的那个函数
		}
	}

}

func Test() {

	var MiGongArr [10][10]int
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			MiGongArr[0][j] = 1
			MiGongArr[9][j] = 1
			MiGongArr[i][0] = 1
			MiGongArr[i][9] = 1
			MiGongArr[3][1] = 1
			MiGongArr[3][2] = 1
			MiGongArr[5][4] = 1
			MiGongArr[6][3] = 1
		}
		fmt.Println(MiGongArr[i])
	}
	fmt.Println()
	MiGong(&MiGongArr, 1, 1)
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
		}
		fmt.Println(MiGongArr[i])
	}

}
