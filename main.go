package main

import (
	"fmt"
)

func countFactor(data int) (res int, err error) {

	for i := 1; i < data; i++ {
		count := 0
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				count++
			}

			if count > 6 {
				break
			}
		}

		if count == 6 {
			res++
		}
	}

	return res, nil
}

func main() {
	count, err := countFactor(134217728)
	fmt.Println(count)
	fmt.Println(err)
}
