package main

import (
	"fmt"
)

func main() {

	for j := 65; j <= 90; j++ {
		fmt.Println(j)

		for i := 1; i <= 3; i++ {

			fmt.Println("\t%#U\n", j)
		}
	}

}
