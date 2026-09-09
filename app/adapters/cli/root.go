package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/shaninalex/lumna/app/cmd/commands"
	"gitlab.com/shaninalex/lumna/app/cmd/commands/board"
	"gitlab.com/shaninalex/lumna/app/cmd/commands/column"
	"gitlab.com/shaninalex/lumna/app/cmd/commands/identity"
	"gitlab.com/shaninalex/lumna/app/cmd/commands/project"
	"gitlab.com/shaninalex/lumna/app/cmd/commands/workspace"
)

func NewRootCmd() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use: "lumna",
	}

	cmd.AddCommand(commands.NewRootServeCommand())
	cmd.AddCommand(commands.NewMigrateRootCmd())
	cmd.AddCommand(commands.NewStatusRootCmd())
	cmd.AddCommand(commands.NewImportRootCmd())
	cmd.AddCommand(commands.NewExportRootCmd())
	cmd.AddCommand(commands.NewEmailRootCmd())

	cmd.AddCommand(workspace.NewWorkspaceRootCmd())
	cmd.AddCommand(identity.NewIdentitiesRootCmd())
	cmd.AddCommand(project.NewProjectsRootCmd())
	cmd.AddCommand(board.NewBoardsRootCmd())
	cmd.AddCommand(column.NewColumnRootCmd())

	cmd.PersistentFlags().String("config", "", "Configuration path. Required.")
	_ = cmd.MarkPersistentFlagRequired("config")
	return cmd
}

// Execute run application
func Execute() int {
	c := NewRootCmd()

	if err := c.Execute(); err != nil {
		fmt.Println(err)
		return 1
	}
	return 0
}
