package main

import (
	"fmt"
	"time"
)

func main() {
	t:=make(map[string]int)
	change(t)
    fmt.Println("请输入change进行转换")
	var c string
	fmt.Scan(&c)
	if c=="change" {
		for k:=range t {
			if _, ok := t[k];ok{
			fmt.Println("input ok")
			}else{
				fmt.Println("输入有误")
			}
		}
		for v:=range t {
			V,_:=time.Parse("2006-01-02 15:04:05",v)
			fmt.Println(V)
		}
	}
}

func change(m map[string]int) {
	fmt.Println("请输入5个时间戳")
	var te1, te2, te3, te4, te5 int64
	fmt.Scan(&te1, &te2, &te3, &te4, &te5)
	t1 := time.Unix(te1, 0)
	t2 := time.Unix(te2, 0)
	t3 := time.Unix(te3, 0)
	t4 := time.Unix(te4, 0)
	t5 := time.Unix(te5, 0)
	T1 := t1.Format("2006-01-02 15:04:05")
	T2 := t2.Format("2006-01-02 15:04:05")
	T3 := t3.Format("2006-01-02 15:04:05")
	T4 := t4.Format("2006-01-02 15:04:05")
	T5 := t5.Format("2006-01-02 15:04:05")
	m[T1] = 1
	m[T2] = 2
	m[T3] = 3
	m[T4] = 4
	m[T5] = 5
}
