package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(testCmd)
}

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Test command",
	Long:  "This is a test command",
	Run: func(cmd *cobra.Command, args []string) {
		println("Hello from test command")
	},
}
