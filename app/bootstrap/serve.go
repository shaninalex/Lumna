package bootstrap

import (
	"fmt"

	"github.com/spf13/cobra"
)

func serveCmd(app *App) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Run http server",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("serve")
		},
	}

	return cmd
}
