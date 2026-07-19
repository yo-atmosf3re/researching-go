package practice

import (
	"context"
	"errors"
	"researching-go/pkg/logger"
	"time"
)

type Metric struct {
	Name  string
	Value float64
}

type Collector interface {
	Collect() (Metric, error)
}

type CpuCollector struct{}

func (c CpuCollector) Collect() (Metric, error) {
	time.Sleep(200 * time.Millisecond)
	logger.Ptc("goroutine is completed")
	return Metric{Name: "cpu_usage", Value: 42.5}, nil
}

type RamCollector struct{}

func (r RamCollector) Collect() (Metric, error) {
	time.Sleep(3 * time.Second) // !!! so long time
	logger.Ptc("goroutine is completed")
	return Metric{Name: "ram_usage", Value: 80.1}, nil
}

type DiskCollector struct{}

func (d DiskCollector) Collect() (Metric, error) {
	time.Sleep(100 * time.Millisecond)
	logger.Ptc("goroutine is completed")
	return Metric{}, errors.New("disk is not responding") // error
}

func RememberingMediumPlus() {
	collectors := []Collector{
		CpuCollector{},
		RamCollector{},
		DiskCollector{},
	}

	logsCh := make(chan []Metric)
	metricCh := make(chan Metric, len(collectors))
	var batch []Metric
	ctx, cancel := context.WithCancel(context.Background())

	for _, collector := range collectors {
		select {
		case <-ctx.Done():
			logger.Ptc("context is done")
			return
		default:
			go func(c Collector) {
				metric, err := c.Collect()
				if err != nil {
					return
				}
				metricCh <- metric
			}(collector)
		}
	}

	timeout := time.After(1 * time.Second)

	for {
		select {
		case metric := <-metricCh:
			batch = append(batch, metric)
		case <-timeout:
			goto endOfLoop
		}
	}
endOfLoop:

	defer cancel()
	go func() {
		logsCh <- batch
		close(logsCh)
		close(metricCh)
	}()

	for log := range logsCh {
		logger.Ptc(log)
	}
}
