package util

import (
	"fmt"
	"log"
	"runtime"
)

func ByteCountSI(b uint64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB",
		float64(b)/float64(div), "kMGTPE"[exp])
}

func GetRuntimeMemStats() runtime.MemStats {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m
}

func LogRuntimeMemStats() {
	m := GetRuntimeMemStats()

	logLine := ""
	logLine += fmt.Sprintf("Alloc = %v MiB", ByteCountSI(m.Alloc))
	logLine += fmt.Sprintf("\tTotalAlloc = %v MiB", ByteCountSI(m.TotalAlloc))
	logLine += fmt.Sprintf("\tSys = %v MiB", ByteCountSI(m.Sys))
	logLine += fmt.Sprintf("\tNumGC = %v\n", m.NumGC)

	log.Println("[MEM STATS] ", logLine)
}
