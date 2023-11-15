package utils

import (
	"fmt"
	"time"
)

func MeasureTime(process string) func() {
	fmt.Printf("start %s", process)
	start := time.Now()

	return func() {
		fmt.Printf("Task %s took: %v \n", process, time.Since(start))
	}

}
