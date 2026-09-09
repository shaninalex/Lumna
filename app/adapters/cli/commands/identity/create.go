package identity

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/pkg/config"
	"gitlab.com/shaninalex/lumna/app/pkg/utils"
	"gitlab.com/shaninalex/lumna/app/services/auth"
	"gitlab.com/shaninalex/lumna/app/services/persistence"
	"gorm.io/gorm"
)

func NewIdentitiesCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [email] [fullname] [password] [active] [workspace_id]",
		Short: "Create identities",
		Long:  "Require 5 arguments - email, fullname, password, active, workspace id",
		Run: func(cmd *cobra.Command, args []string) {
			flags := getIdentitiesCreateCmd(cmd)
			ctx, appCancel := context.WithCancel(context.Background())
			defer appCancel()
			cnf := config.ProvideConfig(flags.configPath)()
			db := persistence.ProvideDB(cnf)

			user := models.Identity{
				FullName: flags.fullName,
				Email:    flags.email,
				Active:   flags.active,
			}

			if err := gorm.G[models.Identity](db).Create(ctx, &user); err != nil {
				panic(err)
			}

			pwdHash, err := auth.CreatePasswordHash(flags.password)
			if err != nil {
				log.Fatal(err)
			}

			credential := models.Credential{
				IdentityID:     user.ID,
				Provider:       "local",
				ProviderUserID: utils.Pointer(fmt.Sprintf("%d", user.ID)),
				Email:          &user.Email,
				PasswordHash:   utils.Pointer(pwdHash),
			}
			if err := gorm.G[models.Credential](db).Create(ctx, &credential); err != nil {
				panic(err)
			}
			identityWorkspace := models.IdentityWorkspace{
				IdentityId:  user.ID,
				WorkspaceId: flags.workspaceId,
			}
			if err := gorm.G[models.IdentityWorkspace](db).Create(ctx, &identityWorkspace); err != nil {
				panic(err)
			}

			log.Printf("User: %s, id: %d, credentials: %d, belongs to workspace id: %d\n", user.Email, user.ID, credential.ID, identityWorkspace.WorkspaceId)
		},
	}

	cmd.PersistentFlags().String("email", "", "Identity email")
	cmd.PersistentFlags().String("full_name", "", "Identity full name")
	cmd.PersistentFlags().String("password", "", "Identity password")
	cmd.PersistentFlags().Bool("active", true, "Identity active")
	cmd.PersistentFlags().Int("workspace_id", 0, "Identity workspace id")

	return cmd
}

type flags struct {
	configPath  string
	email       string
	fullName    string
	password    string
	active      bool
	workspaceId int
}

func getIdentitiesCreateCmd(cmd *cobra.Command) flags {
	configPath, err := cmd.Flags().GetString("config")
	if err != nil {
		panic(err)
	}

	email, err := cmd.Flags().GetString("email")
	if err != nil {
		panic(err)
	}

	fullName, err := cmd.Flags().GetString("full_name")
	if err != nil {
		panic(err)
	}

	password, err := cmd.Flags().GetString("password")
	if err != nil {
		panic(err)
	}

	active, err := cmd.Flags().GetBool("active")
	if err != nil {
		panic(err)
	}

	workspaceId, err := cmd.Flags().GetInt("workspace_id")
	if err != nil {
		panic(err)
	}

	return flags{
		configPath:  configPath,
		email:       email,
		fullName:    fullName,
		password:    password,
		active:      active,
		workspaceId: workspaceId,
	}
}
