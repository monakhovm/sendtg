package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const appVersion = "1.0.0"

var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Shows version",
	Long:    "Displays the current version of the sendtg.",
	Example: "sendtg version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("sendtg v%s\n", appVersion)
	},
}
