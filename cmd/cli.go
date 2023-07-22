package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	cobra.OnInitialize()
}

// Execute is the entry point for the CLI.
func Execute() {
	if err := GetRootCommand().Execute(); err != nil {
		fmt.Println(err)
	}
}
