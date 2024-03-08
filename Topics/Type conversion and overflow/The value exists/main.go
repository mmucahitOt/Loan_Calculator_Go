package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)

	result := number != 0

	fmt.Println(result)
}
