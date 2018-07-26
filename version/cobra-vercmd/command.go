package vercmd

import (
	"github.com/recallsong/go-utils/version"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command for cobra
var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Long:  `Print version information.`,
	Run: func(cmd *cobra.Command, args []string) {
		version.Print()
	},
}

func AddTo(parent *cobra.Command) {
	parent.AddCommand(VersionCmd)
}
