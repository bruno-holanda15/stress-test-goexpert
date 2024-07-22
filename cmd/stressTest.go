/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	url         string
	qtyRequests int8
	concurrency int8
)

// stressTestCmd represents the stressTest command
var stressTestCmd = &cobra.Command{
	Use:   "stress-test",
	Short: "Execute stress test",
	Long: `This commands must execute a stress test to a specific URL:

	If any URL is not passed, the test will assume http://google.com.br with 5 requests and 5 concurrency
	You also can pass the number of requests and concurrency to execute the stress test proposed 
	by Challenge from FullCycle Goexpert.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("stressTest called")
		fmt.Println("URL:", url, "Number of requests:", qtyRequests, "Concurrency:", concurrency)
	},
}

func init() {
	rootCmd.AddCommand(stressTestCmd)
	stressTestCmd.Flags().StringVarP(&url, "url", "u", "http://google.com.br", "URL to execute test")
	stressTestCmd.Flags().Int8VarP(&qtyRequests, "qtyRequests", "r", 5, "Number of requests")
	stressTestCmd.Flags().Int8VarP(&concurrency, "concurrency", "c", 5, "Concurrency to execute test")
}
