package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna"
	"gitlab.com/shaninalex/lumna/app/api/controllers/auth"
	"gitlab.com/shaninalex/lumna/app/api/controllers/board"
	"gitlab.com/shaninalex/lumna/app/api/controllers/column"
	"gitlab.com/shaninalex/lumna/app/api/controllers/invitation"
	"gitlab.com/shaninalex/lumna/app/api/controllers/project"
	"gitlab.com/shaninalex/lumna/app/api/controllers/task"
	"gitlab.com/shaninalex/lumna/app/api/controllers/user"
	"gitlab.com/shaninalex/lumna/app/api/controllers/workspace"
	"gitlab.com/shaninalex/lumna/app/api/middlewares"
	"gitlab.com/shaninalex/lumna/app/pkg/config"
	"gitlab.com/shaninalex/lumna/app/web"
	"gitlab.com/shaninalex/lumna/app/web/docs"
	"gitlab.com/shaninalex/lumna/app/web/setup"
	"go.uber.org/dig"
)

func ProvideRouter(conf *config.Config) *gin.Engine {
	//router := gin.New()
	//if conf.Env() != "development" {
	//	gin.SetMode(gin.ReleaseMode)
	//}

	//router.RedirectTrailingSlash = false
	//router.RedirectFixedPath = false

	router := gin.Default()
	router.GET("/_health", HealthRoute)

	return router
}

func HealthRoute(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"version": lumna.Version,
	})
}

type ApiDeps struct {
	dig.In

	AuthController       *auth.AuthController
	BoardController      *board.BoardController
	StatusController     *column.Controller
	ProjectController    *project.ProjectController
	TaskController       *task.TaskController
	UserController       *user.UserController
	InvitationController *invitation.InvitationController
	WorkspaceController  *workspace.WorkspaceController

	AuthMiddleware middlewares.AuthMiddleware
}

func NewApi(deps ApiDeps, config *config.Config) *gin.Engine {
	// base API middlewares
	router := ProvideRouter(config)
	router.Use(gin.Recovery()) // TODO: write your own recovery middleware
	router.Use(middlewares.LoggingMiddleware())
	router.Use(middlewares.CORSMiddleware())

	setup.RegisterSetupRoute(router, config)
	docs.RegisterDocsRoute(router)
	web.RegisterEmbedRoute(router)

	authGroup := router.Group("/api/v1/auth")
	deps.AuthController.Register(authGroup)

	private := router.Group("/api/v1")
	private.Use(gin.HandlerFunc(deps.AuthMiddleware))

	deps.BoardController.Register(private)
	deps.StatusController.Register(private)
	deps.ProjectController.Register(private)
	deps.TaskController.Register(private)
	deps.UserController.Register(private)
	deps.InvitationController.Register(private)
	deps.WorkspaceController.Register(private)

	return router
}
