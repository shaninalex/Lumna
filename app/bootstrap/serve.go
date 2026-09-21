package bootstrap

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
	api "gitlab.com/shaninalex/lumna/app/adapters/http"
	"gitlab.com/shaninalex/lumna/app/adapters/webui"
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

			// router and routes registration
			router := gin.Default()
			if cfg.SetupEnabled() {
				webui.RegisterSetupRoute(resolve, router)
			}
			if cfg.EmbedSPA() {
				webui.RegisterSPA(router)
			}

			webui.RegisterDocsRoute(router)
			api.RegisterWebsocketRoute(resolve, router)
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

			quit := make(chan os.Signal, 1)
			signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
			<-quit

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			if err := srv.Shutdown(ctx); err != nil {
				log.Fatal("Server forced to shutdown:", err)
			}

			log.Println("Server exiting")
		},
	}

	return cmd
}
