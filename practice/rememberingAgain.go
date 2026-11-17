package practice

import (
	"context"
	"errors"
	"fmt"
	"researching-go/pkg/logger"
	"sync"
	"time"
)

func worker(ch chan string) {
	time.Sleep(100 * time.Millisecond)
	ch <- "готово"
}

func startWorker() {
	ch := make(chan string)
	go worker(ch)

	res := <-ch
	logger.Ptc("result: ", res)
}

func simpleWgWait() {
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(2)
		go func(num int) {
			defer wg.Done()
			logger.Ptc("goroutine: ", num)
		}(i)
	}
	wg.Wait()
	logger.Ptc("finished")
}

func chTable() {
	ch := make(chan int, 2)
	ch <- 10
	ch <- 20
	close(ch)

	v1, ok1 := <-ch
	v2, ok2 := <-ch
	v3, ok3 := <-ch
	logger.Ptc(v1, ok1)
	logger.Ptc(v2, ok2)
	logger.Ptc(v3, ok3)
}

type Item struct {
	Name  string
	Price int
}

func stringExample() {
	cart := []Item{
		{Name: "Хлеб", Price: 50},
		{Name: "Молоко", Price: 100},
	}

	for i, _ := range cart {
		cart[i].Price += cart[i].Price
	}

	logger.Ptc("1: ", cart[0].Price)
	word := "GoРулит"
	logger.Ptc("2: ", len(word), len([]rune(word)))
}

var ErrNotFound = errors.New("not found")

func getUserFromDB() error {
	return ErrNotFound
}

func getUserService() error {
	err := getUserFromDB()
	if err != nil {
		return fmt.Errorf("service failed: %v", err)
	}
	return nil
}

func errorExample() {
	err := getUserService()
	fmt.Println("Is ErrNotFound:", errors.Is(err, ErrNotFound))
}

func fetchData(ctx context.Context) {
	select {
	case <-time.After(500 * time.Millisecond):
		logger.Ptc("Данные получены")
	case <-ctx.Done():
		logger.Ptc("Отмена: ", ctx.Err())
	}
}

func contextExample() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	fetchData(ctx)
}

func printLength(data any) {
	s, ok := data.(string)
	if !ok {
		logger.Ptc("Это не строка")
		return
	}
	logger.Ptc("Длина строки: ", len(s))
}

func anyTypeExample() {
	printLength("hello")
	printLength(12345)

	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	ch1 <- "первый"
	ch2 <- "второй"

	select {
	case m1 := <-ch1:
		logger.Ptc(m1)
	case m2 := <-ch2:
		logger.Ptc(m2)
	}
}

type Response[T any] struct {
	Success bool `json:"success"`
	Data    T    `json:"data"`
}

func Contains[T comparable](slice []T, target T) bool {
	for _, item := range slice {
		if item == target {
			return true
		}
	}
	return false
}

type User struct {
	Name string
}

func genericExample() {
	res := Response[string]{Success: true, Data: "токен_доступа"}
	logger.Ptc(res.Data)
	userResp := Response[User]{Success: true, Data: User{
		Name: "Alex",
	}}
	logger.Ptc(userResp.Data)
	logger.Ptc(Contains([]int{10, 20, 30}, 20))
}

type ctxKey string

const roleKey ctxKey = "role"

func childWorker(ctx context.Context) {
	<-ctx.Done()
	fmt.Println("Child остановлен, причина: ", ctx.Err())
}

func contextExampleAgain() {
	rootCtx, cancelRoot := context.WithCancel(context.Background())
	authCtx := context.WithValue(rootCtx, roleKey, "admin")

	go childWorker(authCtx)

	time.Sleep(100 * time.Millisecond)
	cancelRoot()
	time.Sleep(50 * time.Millisecond)

	role, ok := rootCtx.Value(roleKey).(string)
	fmt.Println("Role a rootCtx: ", role, ok)
}

var ErrInvalidTask = errors.New("invalid task")

type Stats struct {
	mu        sync.Mutex
	Processed int
	History   []string
}

func (s Stats) Record(task string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Processed++
	s.History = append(s.History, task)
}

func processTask(ctx context.Context, task string) error {
	if task == "" {
		return fmt.Errorf("validation: %w", ErrInvalidTask)
	}

	select {
	case <-time.After(100 * time.Millisecond):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func hardTask() {
	ctx, cancel := context.WithTimeout(context.Background(), 2000*time.Millisecond)
	defer cancel()

	tasks := []string{"job_1", "", "job_2", "job_3"}
	stats := &Stats{}
	errCh := make(chan error)

	wg := &sync.WaitGroup{}
	for _, t := range tasks {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := processTask(ctx, t)
			if err != nil {
				if errors.Is(err, ErrInvalidTask) {
					logger.Ptc("Пропущена пустая задача")
					return
				}
				errCh <- err
				return
			}
			stats.Record(t)
		}()
	}
	wg.Wait()

	close(errCh)
	for err := range errCh {
		logger.Ptc("Ошибка выполнения: ", err)
	}
}

func RememberingAgain() {
	//startWorker()
	//simpleWgWait()
	//chTable()
	//stringExample()
	//contextExample()
	//anyTypeExample()
	//genericExample()
	//contextExampleAgain()
	hardTask()
}
