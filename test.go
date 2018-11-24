package  main

// 导入标准库或第三方库

import (
	"fmt"
	"time"
)

// 程序入口 main函数
func gofun(){
	fmt.Println("go fun")
}

func main(){

    fmt.Println("Hello world!");
	fmt.Println("first go");
	
	for{
	//	go gofun()
		time.Sleep(time.Duration(2)*time.Second)
	}
}