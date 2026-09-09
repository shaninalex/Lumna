package commands

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"gitlab.com/shaninalex/lumna/app/api"
	"gitlab.com/shaninalex/lumna/app/pkg/config"
	"gitlab.com/shaninalex/lumna/app/repositories"
	"gitlab.com/shaninalex/lumna/app/services/email"
	"gitlab.com/shaninalex/lumna/app/services/observer"
	"gitlab.com/shaninalex/lumna/app/services/persistence"
	"go.uber.org/dig"
)

func NewRootServeCommand() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "serve",
		Short: "Run webserver",
		Args:  cobra.ArbitraryArgs,
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
			_ = c.Provide(observer.ProvideEventBus)
			_ = c.Provide(persistence.ProvideDB)

			_ = api.Module(c)
			_ = repositories.Module(c)
			_ = email.Module(c)

			err = c.Invoke(func(router *gin.Engine, config *config.Config, ctx context.Context, _ *email.EmailQueue) {
				srv := &http.Server{
					Addr:    fmt.Sprintf(":%d", config.Int("serve.port")),
					Handler: router,
				}

				log.Printf("Run server on :%d\n", config.Int("serve.port"))
				go func() {
					if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
						log.Fatalf("listen: %s\n", err)
					}
				}()

				quit := make(chan os.Signal, 1)
				signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
				<-quit

				log.Println("Shutting down server...")
				appCancel()

				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()

				if err := srv.Shutdown(ctx); err != nil {
					log.Fatal("Server forced to shutdown:", err)
				}

				log.Println("Server exiting")
			})
			if err != nil {
				panic(err)
			}
		},
	}

	return cmd
}
