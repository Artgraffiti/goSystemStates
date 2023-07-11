package metrics

import (
	"expvar"
	"math/rand"
	"runtime"

	"github.com/shirou/gopsutil/v3/mem"
)

type Metric map[string]uint64

func Get() (metric Metric, err error) {
	memstatsFunc := expvar.Get("memstats").(expvar.Func)
	memstats := memstatsFunc().(runtime.MemStats)
	v, err := mem.VirtualMemory()
	if err != nil {
		return
	}

	metric = make(map[string]uint64)

	metric["BuckHashSys"] = memstats.BuckHashSys
	metric["Frees"] = memstats.Frees
	metric["GCCPUFraction"] = uint64(memstats.GCCPUFraction)
	metric["GCSys"] = memstats.GCSys

	metric["HeapAlloc"] = memstats.HeapAlloc
	metric["HeapIdle"] = memstats.HeapIdle
	metric["HeapInuse"] = memstats.HeapInuse
	metric["HeapObjects"] = memstats.HeapObjects
	metric["HeapReleased"] = memstats.HeapReleased
	metric["HeapSys"] = memstats.HeapSys

	metric["LastGC"] = memstats.LastGC
	metric["Lookups"] = memstats.Lookups

	metric["MCacheInuse"] = memstats.MCacheInuse
	metric["MCacheSys"] = memstats.MCacheSys
	metric["MSpanInuse"] = memstats.MSpanInuse
	metric["MSpanSys"] = memstats.MSpanSys

	metric["Mallocs"] = memstats.Mallocs
	metric["NextGC"] = memstats.NextGC
	metric["NumForcedGC"] = uint64(memstats.NumForcedGC)
	metric["NumGC"] = uint64(memstats.NumGC)
	metric["OtherSys"] = memstats.OtherSys
	metric["PauseTotalNs"] = memstats.PauseTotalNs

	metric["StackInuse"] = memstats.StackInuse
	metric["StackSys"] = memstats.StackSys

	metric["Alloc"] = memstats.Alloc
	metric["Sys"] = memstats.Sys
	metric["TotalAlloc"] = memstats.TotalAlloc
	metric["RandomValue"] = rand.Uint64()

	// Memory
	metric["TotalMemory"] = v.Total
	metric["FreeMemory"] = v.Free
	return
}
