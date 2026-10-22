package bootstrap

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	api "gitlab.com/shaninalex/lumna/app/adapters/api"
	"gitlab.com/shaninalex/lumna/app/adapters/web"
	"gitlab.com/shaninalex/lumna/app/adapters/ws"
	"gitlab.com/shaninalex/lumna/app/core"
	"gitlab.com/shaninalex/lumna/app/platform/config"
)

// appRef resolves the application after PersistentPreRunE has built it.
// serve is composition, not an adapter, so it is allowed to see the whole App
// and hand module bridges to the adapters that need them.
type appRef func() *Instance

func serveCmd(resolve core.Resolve, app appRef) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Run http server",
		Run: func(cmd *cobra.Command, args []string) {
			path, _ := cmd.Flags().GetString("config")
			cfg := config.ReadConfig(path)

			ctx, stop := signal.NotifyContext(cmd.Context(), syscall.SIGINT, syscall.SIGTERM)
			defer stop()

			// workers
			var workers sync.WaitGroup
			if on, _ := cmd.Flags().GetBool("email-queue"); on {
				fmt.Println("Email queue started")
				workers.Go(func() {
					app().Workers.EmailDispatcher.Run(ctx)
				})
			}

			// router and routes registration
			router := gin.Default()
			web.RegisterWebRoutes(cfg, resolve, router)
			ws.RegisterWebsocketRoute(app().Realtime, app().Bridges.AuthVerifier, router)
			api.RegisterApiRoutes(resolve, app().Bridges.AuthVerifier, api.Config{
				CORSOrigins:   cfg.CORSOrigins(),
				SecureCookies: cfg.SecureCookies(),
			}, router)

			// server
			srv := &http.Server{
				Addr:    fmt.Sprintf(":%d", cfg.Int("serve.port")),
				Handler: router,
			}

			log.Printf("Run server on :%d\n", cfg.Int("serve.port"))
			go func() {
				if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
					log.Fatalf("listen: %s\n", err)
				}
			}()

			<-ctx.Done()

			shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			if err := srv.Shutdown(shutdownCtx); err != nil {
				log.Println("Server forced to shutdown:", err)
			}

			workersDone := make(chan struct{})
			go func() {
				workers.Wait()
				close(workersDone)
			}()

			select {
			case <-workersDone:
			case <-shutdownCtx.Done():
				log.Println("Workers did not stop in time")
			}

			log.Println("Server exiting")
		},
	}

	cmd.Flags().Bool("email-queue", false, "Send queued emails from this process. Run it in exactly one process at a time")

	return cmd
}
