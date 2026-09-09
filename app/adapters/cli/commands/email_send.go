package commands

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/spf13/cobra"
	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/pkg/config"
	"gitlab.com/shaninalex/lumna/app/repositories"
	"gitlab.com/shaninalex/lumna/app/services/email"
	"gitlab.com/shaninalex/lumna/app/services/persistence"
	"go.uber.org/dig"
)

func NewEmailCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [path to file]",
		Short: "Create board",
		Long:  "Require 1 argument - path to file.",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			c := dig.New()
			configPath, err := cmd.Flags().GetString("config")
			if err != nil {
				panic(err)
			}

			appContext, appCancel := context.WithCancel(context.Background())
			defer appCancel()
			_ = c.Provide(func() context.Context {
				return appContext
			})
			_ = c.Provide(config.ProvideConfig(configPath))
			_ = c.Provide(persistence.ProvideDB)
			_ = repositories.Module(c)
			_ = email.Module(c)

			var eml models.Email
			f, err := os.Open(args[0])
			if err != nil {
				panic(err)
			}

			b, err := io.ReadAll(f)
			if err != nil {
				panic(err)
			}

			if err = json.Unmarshal(b, &eml); err != nil {
				panic(err)
			}
			fmt.Println(eml)

			if err = c.Invoke(func(ctx context.Context, emailRepository repositories.EmailRepository, _ *email.EmailQueue) {
				if err := emailRepository.Create(ctx, &eml); err != nil {
					panic(err)
				}

				if err := waitForEmail(ctx, emailRepository, eml.ID); err != nil {
					panic(err)
				}
			}); err != nil {
				panic(err)
			}
		},
	}

	return cmd
}

func waitForEmail(ctx context.Context, repo repositories.EmailRepository, id int) error {
	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			e, err := repo.Get(ctx, id)
			if err != nil {
				return err
			}
			switch e.Status {
			case models.EmailStatusSuccess, models.EmailStatusError, models.EmailStatusSkipped:
				fmt.Printf("email %d finished with status: %s\n", e.ID, e.Status)
				return nil
			}
		}
	}
}
