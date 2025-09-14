package main

import (
	"fmt"
)

func main() {

	fmt.Println(TestAddition(2, 4))

}

func TestAddition(v1 int, v2 int) int {
	return v1 + v2
}
