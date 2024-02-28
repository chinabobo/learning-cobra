package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Bobo",
	Long:  `All software has versions. This is Bobo's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Bobo Static Site Generator v1.0 -- HEAD")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
