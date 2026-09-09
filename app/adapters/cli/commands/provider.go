package commands

import (
	"context"

	"github.com/spf13/cobra"
)

func ProvideContext(cmd *cobra.Command) context.Context {
	return cmd.Context()
}

func ProvideCobraCmd(cmd *cobra.Command) func() *cobra.Command {
	return func() *cobra.Command {
		return cmd
	}
}
