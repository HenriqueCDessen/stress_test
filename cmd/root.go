package cmd

import (
	"fmt"
	"os"

	"github.com/henriquedessen/stress_test/internal/runner"
	"github.com/spf13/cobra"
)

var (
	url         string
	total       int
	concurrency int
)

var rootCmd = &cobra.Command{
	Use:   "load-tester",
	Short: "CLI para teste de carga",
	Run: func(cmd *cobra.Command, args []string) {
		if err := runner.RunTest(url, total, concurrency); err != nil {
			fmt.Println("Erro:", err)
			os.Exit(1)
		}
	},
}

func Execute() {
	rootCmd.Flags().StringVar(&url, "url", "", "URL do serviço")
	rootCmd.Flags().IntVar(&total, "requests", 1, "Total de requisições")
	rootCmd.Flags().IntVar(&concurrency, "concurrency", 1, "Número de chamadas simultâneas")
	rootCmd.MarkFlagRequired("url")

	rootCmd.Execute()
}
