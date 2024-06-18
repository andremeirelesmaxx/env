package main

import (
	"github.com/maxxcard/env-inject/internal/cmd"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "injectenv",
		Short: "show options",
	}

	rootCmd.AddCommand(cmd.WriteEnvs())

	rootCmd.Execute()
}
