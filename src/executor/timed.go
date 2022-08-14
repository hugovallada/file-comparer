package executor

import (
	"fmt"
	"time"
)

func TimedExecution(fn func()) {
	startTime := time.Now()
	fn()
	fmt.Println("Duration of execution:", time.Since(startTime))
}
