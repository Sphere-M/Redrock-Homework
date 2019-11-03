package main

import (
	"fmt"
)

var (
	myres = make(map[int]int,20)
	chanle=make(chan int,2)
	chanle2=make(chan int,2)
)

func factorial(n int) {
	var res = 1
	for i:=1;i<=n;i++ {
		res*=i

	}
	chanle<- res
	myres[n]=<-chanle
}

func main() {
	for i:=1;i<=20;i++{
		go factorial(i)
	}
	for i,v:=range myres{
		fmt.Printf("myres[%d]=%d\n",i,v)
	}
}

