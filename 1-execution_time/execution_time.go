package execution_time

import (
	"time"
)

func MeasureExecutionTime(f func()) time.Duration {
	start := time.Now()
	f()
	return time.Since(start)
}
