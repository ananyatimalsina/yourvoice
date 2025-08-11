package routes

import (
	"html/template"
	"net/http"
	"runtime"
	"time"
)

var startTime = time.Now()

func StatsHandler(w http.ResponseWriter, r *http.Request) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	uptime := time.Since(startTime)

	// Convert bytes to MB for better readability
	allocMB := float64(m.Alloc) / 1024 / 1024
	totalAllocMB := float64(m.TotalAlloc) / 1024 / 1024
	sysMB := float64(m.Sys) / 1024 / 1024

	tmpl := template.Must(template.ParseFiles("web/templates/stats.html"))
	tmpl.Execute(w, map[string]any{
		"GoVersion":    runtime.Version(),
		"CpuCores":     runtime.NumCPU(),
		"Goroutines":   runtime.NumGoroutine(),
		"Uptime":       uptime.Round(time.Second).String(),
		"CurrentAlloc": allocMB,
		"TotalAlloc":   totalAllocMB,
		"SystemMemory": sysMB,
		"GcCycles":     m.NumGC,
	})
}
