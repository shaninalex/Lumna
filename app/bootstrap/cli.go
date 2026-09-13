package bootstrap

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"gitlab.com/shaninalex/lumna/app/adapters/cli"
	"gitlab.com/shaninalex/lumna/app/core"
	"gitlab.com/shaninalex/lumna/app/core/errs"
	"gitlab.com/shaninalex/lumna/app/platform/config"
)

func RunCLI(assets Assets) int {
	root, cleanup := NewCLI(assets)
	defer cleanup()

	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(*errs.Error)
			if ok {
				fmt.Printf("Error: %s - %s\nExit with: %d\n", err.Code, err.Message, err.Kind)
				return
			}
			fmt.Println("Error: ", r)
		}
	}()

	if err := root.Execute(); err != nil {
		panic(err)
	}
	return 0
}

func NewCLI(assets Assets) (*cobra.Command, func()) {
	var app *Instance

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
	resolve := core.Resolve(func() *core.App { return app.Core })

	root.AddCommand(
		// Composition. 2 very different application entrypoint.
		serveCmd(resolve, func() *Instance { return app }),
		migrateCmd(resolve),

		// regular cli commands
		cli.NewIdentitiesRootCmd(resolve),
		cli.NewWorkspaceRootCmd(resolve),
		cli.NewProjectRootCmd(resolve),
		cli.NewScopeRootCmd(resolve),
		cli.NewStageRootCmd(resolve),
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
