package main

import (
	"fmt"
	"time"
)

func main() {
	x := 0
	for x <= 10 {
		fmt.Println("Value of x is : ", x)
		x++
	}

	y := 10
	for y >= 0 {
		fmt.Println("Value of y is : ", y)
		time.Sleep(time.Second)
		y = y - 1
	}

}
