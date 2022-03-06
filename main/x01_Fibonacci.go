package main

import "fmt"

/*
存储前100项斐波那契数列，并输出
 */



func main(){

	for i:=0;i<10;i++{
		fmt.Println(fibonacci(i))
	}

}

func fibonacci(n int) int{
	if n < 2 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}