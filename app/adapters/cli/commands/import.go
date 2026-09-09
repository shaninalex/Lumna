package commands

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"gitlab.com/shaninalex/lumna/app/pkg/config"
	"gitlab.com/shaninalex/lumna/app/services/persistence"
	"go.uber.org/dig"
)

func NewImportRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "import [path_to_file]",
		Short: "Import db",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			c := dig.New()
			configPath, err := cmd.Flags().GetString("config")
			if err != nil {
				panic(err)
			}

			_ = c.Provide(config.ProvideConfig(configPath))
			_ = c.Provide(persistence.ProvideDB)

			_, err = os.ReadFile(args[0])
			if err != nil {
				log.Fatal(err)
			}

			panic("Not implemented yet")
		},
	}
	return cmd
}
