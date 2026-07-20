package practice

import (
	"context"
	"errors"
	"researching-go/pkg/logger"
	"sync"
	"time"
)

func Ptr[T any](v T) *T {
	return &v
}

type Metric struct {
	Name        string
	FloatValue  *float64
	StringValue *string
}

type Collector interface {
	Collect(ctx context.Context) (Metric, error)
}

type CpuCollector struct{}

func (c CpuCollector) Collect(ctx context.Context) (Metric, error) {
	select {
	case <-time.After(200 * time.Millisecond):
		logger.Ptc("goroutine is completed")
		return Metric{Name: "cpu_usage", FloatValue: Ptr(42.5)}, nil
	case <-ctx.Done():
		return Metric{}, ctx.Err()
	}
}

type RamCollector struct{}

func (r RamCollector) Collect(ctx context.Context) (Metric, error) {
	select {
	case <-time.After(3 * time.Second): // so long time
		logger.Ptc("goroutine is completed")
		return Metric{Name: "ram_usage", FloatValue: Ptr(80.1)}, nil
	case <-ctx.Done():
		return Metric{Name: "ram_usage"}, ctx.Err()
	}
}

type DiskCollector struct{}

func (d DiskCollector) Collect(ctx context.Context) (Metric, error) {
	select {
	case <-time.After(100 * time.Millisecond):
		logger.Ptc("goroutine is completed")
		return Metric{
			Name: "disk_usage",
		}, errors.New("disk is not responded")
	case <-ctx.Done():
		return Metric{}, ctx.Err()
	}
}

func RememberingMediumPlus() {
	start := time.Now()
	collectors := []Collector{
		CpuCollector{},
		RamCollector{},
		DiskCollector{},
	}

	metricCh := make(chan Metric, len(collectors))
	var batch []Metric
	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	wg.Add(len(collectors))
	for _, collector := range collectors {
		go func(c Collector) {
			defer wg.Done()
			metric, err := c.Collect(ctx)
			if err != nil {
				if errors.Is(err, context.DeadlineExceeded) || errors.Is(err, context.Canceled) {
					return // program is not writing error which happened because of ctx
				}
				if metric.Name == "" {
					metric.Name = "unknown_service"
				}
				metric.StringValue = Ptr(err.Error())
			}
			metricCh <- metric
		}(collector)
	}

	wg.Wait()
	close(metricCh)

	for metric := range metricCh {
		batch = append(batch, metric)
	}

	logger.Ptc(batch)
	duration := time.Since(start)
	logger.Ptc(duration.Seconds())
}
