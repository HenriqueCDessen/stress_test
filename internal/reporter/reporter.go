package reporter

import (
	"fmt"
	"math"
	"time"
)

type Result struct {
	Status   int
	Duration time.Duration
	Error    error
}

func Generate(results chan Result) {
	total := 0
	success := 0
	statusCount := make(map[int]int)

	var totalTime time.Duration
	var minTime time.Duration = time.Hour
	var maxTime time.Duration
	var durations []float64

	var errorSamples []error

	for res := range results {
		total++
		if res.Status == 200 {
			success++
		}
		statusCount[res.Status]++
		totalTime += res.Duration
		durations = append(durations, res.Duration.Seconds())

		if res.Duration < minTime {
			minTime = res.Duration
		}
		if res.Duration > maxTime {
			maxTime = res.Duration
		}
		if res.Status == 0 && res.Error != nil {
			errorSamples = append(errorSamples, res.Error)
		}
	}

	avgTime := totalTime / time.Duration(total)

	var stdDev float64
	if len(durations) > 1 {
		var sum float64
		for _, d := range durations {
			diff := d - avgTime.Seconds()
			sum += diff * diff
		}
		stdDev = math.Sqrt(sum / float64(len(durations)))
	}

	fmt.Println("\n--- Relatório ---")
	fmt.Printf("Total de requests:     %d\n", total)
	fmt.Printf("Sucesso (200):         %d\n", success)
	fmt.Printf("Percentual de sucesso: %.2f%%\n", float64(success)/float64(total)*100)
	fmt.Println("Status HTTP:")
	for code, count := range statusCount {
		fmt.Printf("  %d: %d\n", code, count)
	}

	fmt.Printf("\nTempo total:           %v\n", totalTime)
	fmt.Printf("Tempo médio:           %.2fs\n", avgTime.Seconds())
	fmt.Printf("Tempo mínimo:          %.2fs\n", minTime.Seconds())
	fmt.Printf("Tempo máximo:          %.2fs\n", maxTime.Seconds())
	fmt.Printf("Desvio padrão:         %.2fs\n", stdDev)

	if len(errorSamples) > 0 {
		fmt.Println("\n⚠️  Erros de request (status 0):")
		for i, err := range errorSamples {
			fmt.Printf("  [%d] %v\n", i+1, err)
		}
	}
}
