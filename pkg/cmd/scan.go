package cmd

import "github.com/spf13/cobra"

var scanCmd = &cobra.Command{
    Use:   "scan",
    Short: "Scan cluster for policy violations",
    Run: func(cmd *cobra.Command, args []string) {
        // Implementation
    },
}

func init() {
    rootCmd.AddCommand(scanCmd)
}
