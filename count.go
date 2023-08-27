package main

import (
	"fmt"
)

func main() {

	for count := 0; count < 20; count++ {
		if count == 8 {
			continue
		}

		if count == 15 {
			break
		}

		fmt.Println(count)
	}

}
