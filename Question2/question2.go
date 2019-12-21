package main

import (
	"fmt"
	"io"
	"os"
)


func main(){
	file, err := os.Open(`E:\Golang\GoProjects\src\Redrock Web\Question2\Students.txt`)
	if err != nil {
		fmt.Println("文件打开失败：",err)
	}
	defer file.Close()
	buf := make([]byte, 12)
	data := make([]byte, 0)
	for {
		n, err := file.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}
		data = append(data, buf[:n]...)
	}

	fmt.Println(string(data))
}

