package main

import "fmt"

func main() {
Lable1:
	for {
		for a := 0; a < 10; a++ {
			fmt.Println(a)
			if a > 5 {
				break Lable1
			}
		}
	}
	fmt.Println("结束了")
}
