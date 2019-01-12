package monitor

import (
	"encoding/json"
	"runtime"
	"fmt"
)

type Monitor struct {
	Alloc,
	TotalAlloc,
	Sys,
	Mallocs,
	Frees,
	LiveObjects,
	PauseTotalNs uint64

	NumGC        uint32
	NumGoroutine int
}

func PrintMonitor() {
	var m Monitor
	var rtm runtime.MemStats

	// Read full mem stats
	runtime.ReadMemStats(&rtm)

	// Number of goroutines
	m.NumGoroutine = runtime.NumGoroutine()

	// Misc memory stats
	m.Alloc = bToMb(rtm.Alloc)
	m.TotalAlloc = bToMb(rtm.TotalAlloc)
	m.Sys = bToMb(rtm.Sys)
	m.Mallocs = bToMb(rtm.Mallocs)
	m.Frees = bToMb(rtm.Frees)

	// Live objects = Mallocs - Frees
	m.LiveObjects = m.Mallocs - m.Frees

	// GC Stats
	m.PauseTotalNs = rtm.PauseTotalNs
	m.NumGC = rtm.NumGC

	// Just encode to json and print
	b, _ := json.Marshal(m)

	fmt.Println(string(b))

}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
