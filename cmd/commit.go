package cmd

import (

	_ "github.com/fwessels/s3git-go"
	"github.com/spf13/cobra"
)

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Commit the changes in the repository",
	Long: "Commit the changes in the repository",
	Run: func(cmd *cobra.Command, args []string) {

		// TODO: Add commit command
	},
}

func init() {
	RootCmd.AddCommand(commitCmd)

	// Add local message flags
	commitCmd.Flags().StringP("message", "m", "", "Message for the commit")
}
