package practice

import (
	"researching-go/pkg/logger"
	"sync"
)

type stats struct {
	visits map[string]int
	mtx    sync.Mutex // mutes as field can be transferred to struct, it's comfortable
}

func newStats(mtx sync.Mutex) *stats {
	return &stats{
		mtx:    mtx,
		visits: make(map[string]int), // should be always initialized map from make, else will be nil
	}
}

func increment(url string, stats *stats, wg *sync.WaitGroup) {
	stats.mtx.Lock() // map is weak for write/read, mutex is required
	defer stats.mtx.Unlock()
	defer wg.Done()

	stats.visits[url]++ // without check exist this field, we can set value for this field, which don't exist yet. if field is not exist - it be created.
}

func CounterVisitsExample() {
	mtx := sync.Mutex{}
	sts := newStats(mtx)
	wg := &sync.WaitGroup{}
	initValue := 10_00

	wg.Add(initValue)
	for range initValue {
		go increment("111", sts, wg)
	}
	wg.Add(initValue)
	for range initValue {
		go increment("222", sts, wg)
	}
	wg.Add(initValue)
	for range initValue {
		go increment("333", sts, wg)
	}
	wg.Add(initValue * 2)
	for range initValue {
		go increment("444", sts, wg)
		go increment("555", sts, wg)
	}
	wg.Wait()

	mtx.Lock()
	logger.Ptc("total stats: ", sts.visits)
	mtx.Unlock()
}
