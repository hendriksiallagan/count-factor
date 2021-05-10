package main

import (
	"fmt"
	"time"
)

func microTime() float64 {
	loc, _ := time.LoadLocation("UTC")
	now := time.Now().In(loc)
	micSeconds := float64(now.Nanosecond())
	return float64(now.Unix()) + micSeconds
}

func countFactor(data int) (res int, execTime float64, err error) {
	timeStart := microTime()

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

	timeEnd := microTime()
	execTime = timeEnd - timeStart
	return res, execTime, nil
}

func main() {
	count, execTime, err := countFactor(262144)
	fmt.Println(count)
	fmt.Println(execTime)
	fmt.Println(err)
}
