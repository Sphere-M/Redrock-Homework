package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	m=make([]int,0)
	lock sync.Mutex
     	)

func main() {
	for n:=2;n<=10000;n++ {
		go test(n)
	}
	time.Sleep(1 * time.Second)
	k:=len(m)
	fmt.Println(k)
	fmt.Println(m)
}
func test(n int) {
	var t int
	for i := 1; i <= n; i++ {
		if (n % i) == 0 {
			t++
		}
	}
	lock.Lock()
	if t == 2 {
		m = append(m, n)
	}
	lock.Unlock()
}
