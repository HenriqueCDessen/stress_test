package runner

import (
	"sync"

	"github.com/henriquedessen/stress_test/internal/client"
	"github.com/henriquedessen/stress_test/internal/reporter"
)

func RunTest(url string, total, concurrency int) error {
	var wg sync.WaitGroup
	tasks := make(chan int, total)
	results := make(chan reporter.Result, total)

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range tasks {
				status, duration, err := client.DoRequest(url)
				res := reporter.Result{Status: status, Duration: duration}
				if err != nil {
					res.Error = err
				}
				results <- res
			}
		}()
	}

	for i := 0; i < total; i++ {
		tasks <- i
	}
	close(tasks)

	wg.Wait()
	close(results)

	reporter.Generate(results)
	return nil
}
