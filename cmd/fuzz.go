package cmd

import (

	//"time"

	"github.com/momaek/authy/service"
	"github.com/spf13/cobra"
)

// fuzzCmd represents the fuzz command
var fuzzCmd = &cobra.Command{
	Use:   "fuzz",
	Short: "Fuzzy search your otp tokens(case-insensitive)",
	Long: `Fuzzy search your otp tokens(case-insensitive)

First time(or after clean cache) , need your authy main password`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			mode    = ""
			keyword = ""
		)
		if alfredCount != nil && *alfredCount > 0 {
			mode = "alfred"
		}
		if shellCount != nil && *shellCount > 0 {
			mode = "shell"
		}

		if len(args) > 0 {
			keyword = args[0]
		}

		service.NewSearcher(keyword, mode).Search()

	},
}

var alfredCount *int
var shellCount *int

func init() {
	rootCmd.AddCommand(fuzzCmd)
	alfredCount = fuzzCmd.Flags().CountP("alfred", "a", "Specify Output Mode AlfredWorkflow")
	shellCount = fuzzCmd.Flags().CountP("shell", "s", "Specify Output Mode Shell")
}
