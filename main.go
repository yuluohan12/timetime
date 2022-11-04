package main

import (
	"fmt"
	"os"
	"time"
)

var str, str2 string
var T, T1 time.Time
var s, s1 []string

func Select(n int) {
	switch n {
	case 1:
		Function1()
	case 2:
		Function2()
	default:
		fmt.Println("输入错误！")
		os.Exit(0)

	}

}
func Function1() {
	var now time.Time
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("请设置日程提醒的时间（格式为：2006/01/02/15:04:05）")
	fmt.Scan(&str)
	T, err = time.ParseInLocation("2006/01/02/15:04:05", str, loc)
	if err != nil {
		panic(err)
	}
	fmt.Println("请设置日程提醒的内容")
	fmt.Scan(&str)
	s = append(s, str)
	for {
		time.Sleep(500 * time.Millisecond)
		now = time.Now()
		now.Format("2006-01-02 15:04:05")
		fmt.Println("预定成功！")
		if T.Sub(now) < time.Second {
			fmt.Println(s[0])
			s = s[1:]
			break
		}
	}
	return

}
func Function2() {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("请输入设定的时间(格式为：2006-01-02 15：04：05）：")
	fmt.Scanf("%s", &str2)
	T, err = time.ParseInLocation("2006/01/02/15:04:05", str2, loc)
	if err != nil {
		panic(err)
	}
	fmt.Println("请输入日程提醒内容：")
	fmt.Scanf("%s", &str2)
	s = append(s, str2)
	fmt.Println("日程将在每天这个时刻提醒")
	var now time.Time
	for {
		time.Sleep(500 * time.Millisecond)
		now = time.Now()
		if T1.Sub(now) < time.Second {
			fmt.Println(s1[0])
			T1.Add(24 * time.Hour)
		}
	}
	return

}
func main() {
	var n int
	fmt.Println("功能提示如下：")
	fmt.Println("1.单次日程提醒.")
	fmt.Println("2.重复日程提醒.")
	fmt.Println("请输入序号：")
	fmt.Scanf("%d", &n)
	Select(n)

}
