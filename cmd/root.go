package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "myvpn",
	Short: "This is a CLI to manage vpn connections",
	Long:  `This is a CLI to manage vpn connections, add configurations and more`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here

		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func PrintMessage(m string) {
	println(m)
}
