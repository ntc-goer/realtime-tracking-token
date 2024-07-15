package cmd

import (
	"github.com/ntc-goer/parser-exercise/config"
	"github.com/spf13/cobra"
)

var (
	cfg *config.Config
)

func Execute() error {
	// Get configs file
	cfg = config.Load()
	rootCmd := &cobra.Command{
		Use: "Parser",
	}
	rootCmd.AddCommand(serverCmd(), workerCmd())
	err := rootCmd.Execute()
	if err != nil {
		return err
	}
	return nil
}
