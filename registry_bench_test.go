package metrics_test

import (
	"crypto/md5"
	"fmt"
	"go-metrics"
	"testing"
	"time"
)

func BenchmarkGetOrRegisterCounter(b *testing.B) {
	b.ReportAllocs()
	registry := metrics.NewRegistry()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			metrics.GetOrRegisterCounter("handler", registry)
		}
	})
}

func BenchmarkGetOrRegisterCounter_RandomName(b *testing.B) {
	b.ReportAllocs()
	registry := metrics.NewRegistry()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			metrics.GetOrRegisterCounter(fmt.Sprintf("%x", md5.Sum([]byte(time.Now().String()))), registry)
		}
	})
}
