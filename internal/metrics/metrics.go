package metrics

import (
	"expvar"
	"math/rand"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/v3/mem"
)

type Metrics map[string]any

func Get() (metrics Metrics, err error) {
	memstatsFunc := expvar.Get("memstats").(expvar.Func)
	memstats := memstatsFunc().(runtime.MemStats)
	v, err := mem.VirtualMemory()
	if err != nil {
		return
	}

	metrics = make(map[string]any)

	metrics["BuckHashSys"] = memstats.BuckHashSys
	metrics["Frees"] = memstats.Frees
	metrics["GCCPUFraction"] = memstats.GCCPUFraction
	metrics["GCSys"] = memstats.GCSys

	metrics["HeapAlloc"] = memstats.HeapAlloc
	metrics["HeapIdle"] = memstats.HeapIdle
	metrics["HeapInuse"] = memstats.HeapInuse
	metrics["HeapObjects"] = memstats.HeapObjects
	metrics["HeapReleased"] = memstats.HeapReleased
	metrics["HeapSys"] = memstats.HeapSys

	metrics["LastGC"] = memstats.LastGC
	metrics["Lookups"] = memstats.Lookups

	metrics["MCacheInuse"] = memstats.MCacheInuse
	metrics["MCacheSys"] = memstats.MCacheSys
	metrics["MSpanInuse"] = memstats.MSpanInuse
	metrics["MSpanSys"] = memstats.MSpanSys

	metrics["Mallocs"] = memstats.Mallocs
	metrics["NextGC"] = memstats.NextGC
	metrics["NumForcedGC"] = memstats.NumForcedGC
	metrics["NumGC"] = memstats.NumGC
	metrics["OtherSys"] = memstats.OtherSys
	metrics["PauseTotalNs"] = memstats.PauseTotalNs

	metrics["StackInuse"] = memstats.StackInuse
	metrics["StackSys"] = memstats.StackSys

	metrics["Alloc"] = memstats.Alloc
	metrics["Sys"] = memstats.Sys
	metrics["TotalAlloc"] = memstats.TotalAlloc
	rand.Seed(time.Now().UnixNano())
	metrics["RandomValue"] = rand.Uint64()

	// Memory
	metrics["TotalMemory"] = v.Total
	metrics["FreeMemory"] = v.Free
	return
}
