package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of Mugcake",
	Long:  `Project is currently in drafting stages. I'm still thinking about what to do.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`
			If you're reading this there's nothing here yet... 
			But you're more than welcome to stick around and see how my project goes!

			v0.0.0`)
	},
}
