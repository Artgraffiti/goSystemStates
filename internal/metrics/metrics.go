package metrics

import (
	"expvar"
	"math/rand"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/v3/mem"

	pb "GSS/proto"
)

type MetricStorage struct {
	Timestamp   int64              `json:"timestamp"`
	Uint32Data  map[string]uint32  `json:"uint32"`
	Uint64Data  map[string]uint64  `json:"uint64"`
	Float64Data map[string]float64 `json:"float64"`
}

func Get() (metric MetricStorage, err error) {
	rand.Seed(time.Now().UnixNano())
	memstatsFunc := expvar.Get("memstats").(expvar.Func)
	memstats := memstatsFunc().(runtime.MemStats)
	v, err := mem.VirtualMemory()
	if err != nil {
		return
	}

	metric = MetricStorage{
		Timestamp: time.Now().Unix(),
		Uint32Data: map[string]uint32{
			"NumForcedGC": memstats.NumForcedGC,
			"NumGC":       memstats.NumGC,
		},
		Uint64Data: map[string]uint64{
			"BuckHashSys": memstats.BuckHashSys,
			"Frees":       memstats.Frees,
			"GCSys":       memstats.GCSys,

			"HeapAlloc":    memstats.HeapAlloc,
			"HeapIdle":     memstats.HeapIdle,
			"HeapInuse":    memstats.HeapInuse,
			"HeapObjects":  memstats.HeapObjects,
			"HeapReleased": memstats.HeapReleased,
			"HeapSys":      memstats.HeapSys,

			"LastGC":  memstats.LastGC,
			"Lookups": memstats.Lookups,

			"MCacheInuse": memstats.MCacheInuse,
			"MCacheSys":   memstats.MCacheSys,
			"MSpanInuse":  memstats.MSpanInuse,
			"MSpanSys":    memstats.MSpanSys,

			"Mallocs":      memstats.Mallocs,
			"NextGC":       memstats.NextGC,
			"OtherSys":     memstats.OtherSys,
			"PauseTotalNs": memstats.PauseTotalNs,

			"StackInuse": memstats.StackInuse,
			"StackSys":   memstats.StackSys,

			"Alloc":       memstats.Alloc,
			"Sys":         memstats.Sys,
			"TotalAlloc":  memstats.TotalAlloc,
			"RandomValue": rand.Uint64(),

			// Memory
			"TotalMemory": v.Total,
			"FreeMemory":  v.Free,
		},
		Float64Data: map[string]float64{
			"GCCPUFraction": memstats.GCCPUFraction,
		},
	}
	return
}

func ConvertMetricStorageToProto(mStorage *MetricStorage) (proto *pb.MetricStorage) {
	proto = &pb.MetricStorage{
		Timestamp:   mStorage.Timestamp,
		Uint32Data:  mStorage.Uint32Data,
		Uint64Data:  mStorage.Uint64Data,
		Float64Data: mStorage.Float64Data,
	}
	return
}

func ConvertProtoToMetricStorage(proto *pb.MetricStorage) (mStorage *MetricStorage) {
	mStorage = &MetricStorage{
		Timestamp:   proto.Timestamp,
		Uint32Data:  proto.Uint32Data,
		Uint64Data:  proto.Uint64Data,
		Float64Data: proto.Float64Data,
	}
	return
}
