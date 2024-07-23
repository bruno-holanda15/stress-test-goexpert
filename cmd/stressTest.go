/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

var (
	url         string
	qtyRequests int
	concurrency int
)

// stressTestCmd represents the stressTest command
var stressTestCmd = &cobra.Command{
	Use:   "run",
	Short: "Execute stress test",
	Long: `This commands must execute a stress test to a specific URL:

	If any URL is not passed, the test will assume http://google.com.br with 5 requests and 5 concurrency
	You also can pass the number of requests and concurrency to execute the stress test proposed 
	by Challenge from FullCycle Goexpert.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("stressTest called")
		startTime := time.Now()

		var wg sync.WaitGroup
		concurrencyLimiter := make(chan struct{}, concurrency)

		report := make(map[int]int)
		var mu sync.Mutex
		var counter int

		for range qtyRequests {
			wg.Add(1)

			concurrencyLimiter <- struct{}{}
			go makeReq(url, report, &mu, &wg, concurrencyLimiter)
			counter++
		}

		wg.Wait()

		timeSpent := time.Since(startTime)

		fmt.Println("Tempo de execução em segundos:", timeSpent.Seconds())
		fmt.Println("Quantidade de requests executadas:", counter)
		fmt.Println("Requests executadas com sucesso (status code 200):", report[200])
		for i, v := range report {
			fmt.Println("Status Code:", i, "quantidade:", v)
		}
	},
}

func makeReq(url string, report map[int]int, mu *sync.Mutex, wg *sync.WaitGroup, concurrencyLimiter chan struct{}) {
	defer wg.Done()

	res, _ := http.Get(url)
	mu.Lock()
	if _, exists := report[res.StatusCode]; !exists {
		report[res.StatusCode] = 1
	} else {
		report[res.StatusCode]++
	}
	mu.Unlock()

	<-concurrencyLimiter
}

func init() {
	rootCmd.AddCommand(stressTestCmd)
	stressTestCmd.Flags().StringVarP(&url, "url", "u", "http://google.com.br", "URL to execute test")
	stressTestCmd.Flags().IntVarP(&qtyRequests, "qtyRequests", "r", 100, "Number of requests")
	stressTestCmd.Flags().IntVarP(&concurrency, "concurrency", "c", 5, "Concurrency to execute test")
}
