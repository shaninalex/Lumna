package commands

import (
	"github.com/spf13/cobra"
	"gitlab.com/shaninalex/lumna/app/pkg/config"
	"gitlab.com/shaninalex/lumna/app/services/persistence"
	"go.uber.org/dig"
)

func NewExportRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "export [path_to_file]",
		Short: "Export db",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			c := dig.New()
			configPath, err := cmd.Flags().GetString("config")
			if err != nil {
				panic(err)
			}

			_ = c.Provide(config.ProvideConfig(configPath))
			_ = c.Provide(persistence.ProvideDB)
		},
	}
	return cmd
}
