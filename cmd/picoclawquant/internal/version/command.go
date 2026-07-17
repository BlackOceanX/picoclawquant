package version

import (
	"github.com/spf13/cobra"

	"github.com/BlackOceanX/picoclawquant/cmd/picoclawquant/internal"
	"github.com/BlackOceanX/picoclawquant/cmd/picoclawquant/internal/cliui"
	"github.com/BlackOceanX/picoclawquant/pkg/config"
)

func NewVersionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "version",
		Aliases: []string{"v"},
		Short:   "Show version information",
		Run: func(_ *cobra.Command, _ []string) {
			printVersion()
		},
	}

	return cmd
}

func printVersion() {
	build, goVer := config.FormatBuildInfo()
	cliui.PrintVersion(internal.Logo, "picoclaw "+config.FormatVersion(), build, goVer)
}
