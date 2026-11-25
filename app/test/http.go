package test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/adapters/httpx"
	"gitlab.com/shaninalex/lumna/app/core"
	"gitlab.com/shaninalex/lumna/app/core/actor"
	"gitlab.com/shaninalex/lumna/app/core/bus"
)

// NewApp returns an App with empty buses and no middleware.
func NewApp() *core.App {
	return &core.App{
		Commands: bus.NewCommandBus(),
		Queries:  bus.NewQueryBus(),
		Events:   bus.NewEventBus(),
	}
}

// Resolve wraps app in the core.Resolve that controllers take.
func Resolve(app *core.App) core.Resolve {
	return func() *core.App { return app }
}

// NewEngine returns a gin engine in test mode with no middleware. There is
// deliberately no gin.Recovery: Serve turns a handler panic into a test failure.
func NewEngine() *gin.Engine {
	gin.SetMode(gin.TestMode)
	return gin.New()
}

// Calls records what a fake module received.
type Calls struct {
	List []any
}

// RecordCommand registers a fake handler for command C that appends every
// command it gets to calls and returns result.
func RecordCommand[C any, R any](t *testing.T, app *core.App, calls *Calls, result R) {
	t.Helper()
	err := bus.RegisterCommand(app.Commands, func(_ context.Context, cmd C) (R, error) {
		calls.List = append(calls.List, cmd)
		return result, nil
	})
	if err != nil {
		t.Fatal(err)
	}
}

// Response is the decoded API envelope with the data left raw.
type Response = httpx.APIResponse[json.RawMessage]

// RequestOption changes a request before Serve sends it.
type RequestOption func(*http.Request) *http.Request

// WithActor puts a in the request context, the way AuthMiddleware does.
func WithActor(a actor.Actor) RequestOption {
	return func(r *http.Request) *http.Request {
		return r.WithContext(actor.With(r.Context(), a))
	}
}

// Serve sends one JSON request to h and decodes the response envelope.
// A panic in a handler fails the test.
func Serve(t *testing.T, h http.Handler, method, path, body string, opts ...RequestOption) (*httptest.ResponseRecorder, Response) {
	t.Helper()
	req := httptest.NewRequest(method, path, strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	for _, opt := range opts {
		req = opt(req)
	}
	w := httptest.NewRecorder()

	func() {
		defer func() {
			if p := recover(); p != nil {
				t.Fatalf("%s %s panicked: %v", method, path, p)
			}
		}()
		h.ServeHTTP(w, req)
	}()

	var resp Response
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("%s %s: response is not JSON: %v\n%s", method, path, err, w.Body.String())
	}
	return w, resp
}
