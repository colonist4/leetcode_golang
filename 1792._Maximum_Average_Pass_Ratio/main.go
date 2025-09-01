package main

import (
	"fmt"
)

func main() {
	result := maxAverageRatio(
		[][]int{
			{1, 2},
			{3, 5},
			{2, 2},
		},
		2,
	)

	fmt.Println(result)
}
