package identity

import (
	"github.com/spf13/cobra"
	"gitlab.com/shaninalex/lumna/app/bootstrap"
	"gitlab.com/shaninalex/lumna/app/platform/config"
)

func NewIdentitiesListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List identities",
		Run: func(cmd *cobra.Command, args []string) {
			configPath, err := cmd.Flags().GetString("config")
			if err != nil {
				panic(err)
			}

			cfg := config.ProvideConfig(configPath)
			if _, err := bootstrap.New(cmd.Context(), cfg, bootstrap.Assets{}); err != nil {
				panic(err)
			}

		},
	}

	return cmd
}
