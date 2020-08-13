package cmd

import (
	"fmt"

	"github.com/kmdkuk/gorter/version"
	"github.com/spf13/cobra"
)

func NewVersionCmd() *cobra.Command {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Show version and build revision.",
		Long:  `Show version and build revision.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("version: %s-%s (%s)\n", version.Version, version.Revision, version.BuildDate)
		},
	}
	return versionCmd
}
