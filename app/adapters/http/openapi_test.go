package api

import (
	"os"
	"reflect"
	"regexp"
	"sort"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	authctl "gitlab.com/shaninalex/lumna/app/api/controllers/auth"
	"gitlab.com/shaninalex/lumna/app/api/controllers/board"
	"gitlab.com/shaninalex/lumna/app/api/controllers/column"
	"gitlab.com/shaninalex/lumna/app/api/controllers/invitation"
	"gitlab.com/shaninalex/lumna/app/api/controllers/project"
	"gitlab.com/shaninalex/lumna/app/api/controllers/task"
	userctl "gitlab.com/shaninalex/lumna/app/api/controllers/user"
	"gitlab.com/shaninalex/lumna/app/api/controllers/workspace"
	"gopkg.in/yaml.v3"
)

const specPath = "../../resources/openapi/openapi.yaml"

var ginParam = regexp.MustCompile(`:([^/]+)`)

// testRouter mirrors the wiring of NewApi, but with nil services: Register only
// needs the controller receiver, so no handler is ever invoked here.
func testRouter(t *testing.T) *gin.Engine {
	t.Helper()
	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.GET("/_health", HealthRoute)

	authctl.NewAuthContoller(nil, nil).Register(router.Group("/api/v1/auth"))

	private := router.Group("/api/v1")
	board.NewListController(nil, nil, nil).Register(private)
	column.NewStatusController(nil).Register(private)
	project.NewProjectsController(nil, nil).Register(private)
	task.NewTaskController(nil, nil).Register(private)
	userctl.NewUserController(nil, nil).Register(private)
	invitation.NewInvitationController(nil).Register(private)
	workspace.NewWorkspaceController(nil).Register(private)

	// Guard against the wiring above drifting from NewApi: every controller in
	// ApiDeps has to be registered here, or the comparison below is incomplete.
	want := 0
	deps := reflect.TypeOf(ApiDeps{})
	for i := 0; i < deps.NumField(); i++ {
		if strings.HasSuffix(deps.Field(i).Name, "Controller") {
			want++
		}
	}
	if got := 8; got != want {
		t.Fatalf("ApiDeps has %d controllers, testRouter registers %d - add the missing one", want, got)
	}

	return router
}

func specOperations(t *testing.T) map[string]bool {
	t.Helper()

	raw, err := os.ReadFile(specPath)
	if err != nil {
		t.Fatalf("read spec: %v", err)
	}

	var spec struct {
		Paths map[string]map[string]any `yaml:"paths"`
	}
	if err := yaml.Unmarshal(raw, &spec); err != nil {
		t.Fatalf("parse spec: %v", err)
	}

	ops := map[string]bool{}
	for path, item := range spec.Paths {
		for method := range item {
			switch strings.ToUpper(method) {
			case "GET", "POST", "PUT", "PATCH", "DELETE":
				ops[strings.ToUpper(method)+" "+path] = true
			}
		}
	}

	return ops
}

// TestRoutesAreDocumented keeps resources/openapi/openapi.yaml in sync with the
// routes the controllers actually register, in both directions.
func TestRoutesAreDocumented(t *testing.T) {
	spec := specOperations(t)
	routes := map[string]bool{}

	for _, r := range testRouter(t).Routes() {
		routes[r.Method+" "+ginParam.ReplaceAllString(r.Path, "{$1}")] = true
	}

	var undocumented, stale []string

	for op := range routes {
		if !spec[op] {
			undocumented = append(undocumented, op)
		}
	}

	for op := range spec {
		if !routes[op] {
			stale = append(stale, op)
		}
	}

	sort.Strings(undocumented)
	sort.Strings(stale)

	for _, op := range undocumented {
		t.Errorf("route %s is not in the OpenAPI spec", op)
	}
	for _, op := range stale {
		t.Errorf("spec documents %s, but no route serves it", op)
	}
}

// TestPathParamsAreSnakeCase pins the API convention: :board_id, never :boardId.
func TestPathParamsAreSnakeCase(t *testing.T) {
	snake := regexp.MustCompile(`^[a-z][a-z0-9_]*$`)

	for _, r := range testRouter(t).Routes() {
		for _, segment := range strings.Split(r.Path, "/") {
			name, ok := strings.CutPrefix(segment, ":")
			if !ok {
				continue
			}
			if !snake.MatchString(name) {
				t.Errorf("%s %s: path param %q is not snake_case", r.Method, r.Path, name)
			}
		}
	}
}
