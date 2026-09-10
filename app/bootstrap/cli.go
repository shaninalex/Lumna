package bootstrap

import (
	"context"
	"time"

	"github.com/spf13/cobra"
	"gitlab.com/shaninalex/lumna/app/adapters/cli"
	"gitlab.com/shaninalex/lumna/app/core"
	"gitlab.com/shaninalex/lumna/app/platform/config"
)

func RunCLI(assets Assets) int {
	root, cleanup := NewCLI(assets)
	defer cleanup()

	if err := root.Execute(); err != nil {
		panic(err)
		// TODO: get actual code from error
		// return int(errs.KindInternal)
	}
	return 0
}

func NewCLI(assets Assets) (*cobra.Command, func()) {
	var app *App

	root := &cobra.Command{
		Use:           "lumna",
		SilenceUsage:  true,
		SilenceErrors: true,
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			path, _ := cmd.Flags().GetString("config")
			cfg := config.ReadConfig(path)
			var err error
			app, err = New(cmd.Context(), cfg, assets)

			return err
		},
	}

	root.PersistentFlags().String("config", "", "Path to config yaml")
	resolve := cli.Resolve(func() *core.App { return app.Core })

	root.AddCommand(
		// Composition. 2 very different application entrypoints.
		serveCmd(app),
		migrateCmd(app),

		// regular cli commands
		cli.NewIdentitiesRootCmd(resolve),
	)

	cleanup := func() {
		if app != nil {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			_ = app.Close(ctx)
		}
	}
	return root, cleanup
}
