package cmd

import (
	"github.com/spf13/cobra"
)

type packageCommand struct {
	cmd *cobra.Command
}

func newPackageCommand() *packageCommand {
	cmd := &cobra.Command{
		Use:     "package",
		Aliases: []string{"p", "pkg"},
		Short:   "Manage packages",
	}

	cmd.AddCommand(
		newPackageAddCommand().cmd,
		newPackageExportCommand().cmd,
		newPackageImportCommand().cmd,
		newPackageListCommand().cmd,
		newPackageRemoveCommand().cmd,
		newPackageShowCommand().cmd,
	)

	return &packageCommand{cmd: cmd}
}
