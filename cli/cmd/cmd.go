package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mugcake",
	Short: "Hi, welcome to my project!",
	Long: `
	
		There's nothing here yet...
		But you're more than welcome to stick around and see how my project goes!`,
	Run: func(cmd *cobra.Command, args []string) {
		println("Welcome to mugcake!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
