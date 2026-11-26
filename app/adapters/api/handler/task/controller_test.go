package task_test

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/adapters/api/handler/task"
	"gitlab.com/shaninalex/lumna/app/core"
	"gitlab.com/shaninalex/lumna/app/core/bus"
	"gitlab.com/shaninalex/lumna/app/core/errs"
	"gitlab.com/shaninalex/lumna/app/modules/tracker/contract"
	"gitlab.com/shaninalex/lumna/app/test"
)

func newTrackerApp(t *testing.T) (*core.App, *test.Calls) {
	t.Helper()
	app := test.NewApp()
	calls := &test.Calls{}
	test.RecordCommand[contract.WorkItemCreate](t, app, calls, contract.WorkItemView{})
	test.RecordCommand[contract.WorkItemUpdate](t, app, calls, contract.WorkItemView{})
	test.RecordCommand[contract.WorkItemAssign](t, app, calls, true)
	test.RecordCommand[contract.WorkItemDelete](t, app, calls, true)
	test.RecordCommand[contract.WorkItemMove](t, app, calls, contract.WorkItemView{})
	test.RecordCommand[contract.WorkItemTransfer](t, app, calls, contract.WorkItemView{})
	test.RecordCommand[contract.StageMove](t, app, calls, contract.StageView{})
	return app, calls
}

func newRouter(app *core.App) *gin.Engine {
	r := test.NewEngine()
	task.Register(test.Resolve(app), r.Group("/tasks"))
	return r
}

const (
	codeValidation           = "VALIDATION_ERROR"
	codeUnsupportedMediaType = "UNSUPPORTED_MEDIA_TYPE"
)

func TestTaskController(t *testing.T) {
	tests := []struct {
		name       string
		method     string
		path       string
		body       string
		header     map[string]string // overrides, e.g. Content-Type
		wantStatus int
		wantCode   string // first error code; "" means a successful response
		wantCmd    any    // the one command the module must receive; nil means none
	}{
		// POST /tasks
		{
			name: "create: valid payload reaches the module", method: http.MethodPost, path: "/tasks",
			body:       `{"title":"Write tests","body":"for #13","project_id":1,"column_id":2,"board_id":3,"position":1.5}`,
			wantStatus: http.StatusOK,
			wantCmd: contract.WorkItemCreate{
				Title: "Write tests", Description: "for #13", ProjectId: 1, StageId: 2, ScopeId: 3, Position: 1.5,
			},
		},
		{
			name: "create: empty title (issue #13 repro)", method: http.MethodPost, path: "/tasks",
			body: `{"title":""}`, wantStatus: http.StatusBadRequest, wantCode: codeValidation,
		},
		{
			name: "create: missing title", method: http.MethodPost, path: "/tasks",
			body: `{}`, wantStatus: http.StatusBadRequest, wantCode: codeValidation,
		},
		{
			name: "create: malformed JSON", method: http.MethodPost, path: "/tasks",
			body: `{`, wantStatus: http.StatusBadRequest, wantCode: codeValidation,
		},
		{
			name: "create: empty body", method: http.MethodPost, path: "/tasks",
			body: ``, wantStatus: http.StatusBadRequest, wantCode: codeValidation,
		},
		{
			name: "create: text/plain body (issue #16 CSRF repro)", method: http.MethodPost, path: "/tasks",
			body:       `{"title":"Write tests","x":"="}`,
			header:     map[string]string{"Content-Type": "text/plain"},
			wantStatus: http.StatusUnsupportedMediaType, wantCode: codeUnsupportedMediaType,
		},
		{
			name: "create: form-urlencoded body", method: http.MethodPost, path: "/tasks",
			body:       `title=Write+tests`,
			header:     map[string]string{"Content-Type": "application/x-www-form-urlencoded"},
			wantStatus: http.StatusUnsupportedMediaType, wantCode: codeUnsupportedMediaType,
		},
		{
			name: "create: multipart body", method: http.MethodPost, path: "/tasks",
			body:       `--b--`,
			header:     map[string]string{"Content-Type": "multipart/form-data; boundary=b"},
			wantStatus: http.StatusUnsupportedMediaType, wantCode: codeUnsupportedMediaType,
		},
		{
			name: "create: no Content-Type", method: http.MethodPost, path: "/tasks",
			body:       `{"title":"Write tests"}`,
			header:     map[string]string{"Content-Type": ""},
			wantStatus: http.StatusUnsupportedMediaType, wantCode: codeUnsupportedMediaType,
		},
		{
			name: "create: JSON with charset passes", method: http.MethodPost, path: "/tasks",
			body:       `{"title":"Write tests"}`,
			header:     map[string]string{"Content-Type": "application/json; charset=utf-8"},
			wantStatus: http.StatusOK,
			wantCmd:    contract.WorkItemCreate{Title: "Write tests"},
		},
		{
			name: "create: wrong field type", method: http.MethodPost, path: "/tasks",
			body: `{"title":42}`, wantStatus: http.StatusBadRequest, wantCode: codeValidation,
		},

		// PATCH /tasks/:task_id
		{
			name: "update: id comes from the path", method: http.MethodPatch, path: "/tasks/5",
			body: `{"title":"New title","body":"New body"}`, wantStatus: http.StatusOK,
			wantCmd: contract.WorkItemUpdate{WorkItemId: 5, Title: "New title", Description: "New body"},
		},
		{
			name: "update: task_id in the body is ignored", method: http.MethodPatch, path: "/tasks/5",
			body: `{"task_id":7,"title":"New title"}`, wantStatus: http.StatusOK,
			wantCmd: contract.WorkItemUpdate{WorkItemId: 5, Title: "New title"},
		},
		{
			name: "update: non-numeric id", method: http.MethodPatch, path: "/tasks/abc",
			body: `{"title":"New title"}`, wantStatus: http.StatusBadRequest, wantCode: codeValidation,
		},
		{
			name: "update: empty title", method: http.MethodPatch, path: "/tasks/5",
			body: `{"title":""}`, wantStatus: http.StatusBadRequest, wantCode: codeValidation,
		},
		{
			name: "update: text/plain body", method: http.MethodPatch, path: "/tasks/5",
			body:       `{"title":"New title"}`,
			header:     map[string]string{"Content-Type": "text/plain"},
			wantStatus: http.StatusUnsupportedMediaType, wantCode: codeUnsupportedMediaType,
		},
		{
			name: "update: malformed JSON", method: http.MethodPatch, path: "/tasks/5",
			body: `{`, wantStatus: http.StatusBadRequest, wantCode: codeValidation,
		},

		// PATCH /tasks/:task_id/assign
		{
			name: "assign: id comes from the path", method: http.MethodPatch, path: "/tasks/5/assign",
			body: `{"identity_id":9}`, wantStatus: http.StatusOK,
			wantCmd: contract.WorkItemAssign{WorkItemId: 5, IdentityId: 9},
		},
		{
			name: "assign: task_id in the body is ignored", method: http.MethodPatch, path: "/tasks/5/assign",
			body: `{"task_id":7,"identity_id":9}`, wantStatus: http.StatusOK,
			wantCmd: contract.WorkItemAssign{WorkItemId: 5, IdentityId: 9},
		},
		{
			name: "assign: non-numeric id", method: http.MethodPatch, path: "/tasks/abc/assign",
			body: `{"identity_id":9}`, wantStatus: http.StatusBadRequest, wantCode: codeValidation,
		},
		{
			name: "assign: missing identity_id", method: http.MethodPatch, path: "/tasks/5/assign",
			body: `{}`, wantStatus: http.StatusBadRequest, wantCode: codeValidation,
		},

		// DELETE /tasks/:task_id
		{
			name: "delete: id comes from the path", method: http.MethodDelete, path: "/tasks/5",
			wantStatus: http.StatusOK, wantCmd: contract.WorkItemDelete{WorkItemId: 5},
		},
		{
			name: "delete: non-numeric id", method: http.MethodDelete, path: "/tasks/abc",
			wantStatus: http.StatusBadRequest, wantCode: codeValidation,
		},

		// POST /tasks/move
		{
			name: "move: move_task", method: http.MethodPost, path: "/tasks/move",
			body:       `{"action":"move_task","data":{"board_id":3,"task_id":5,"position":2.5}}`,
			wantStatus: http.StatusOK,
			wantCmd:    contract.WorkItemMove{ScopeId: 3, WorkItemId: 5, Rank: 2.5},
		},
		{
			name: "move: move_column", method: http.MethodPost, path: "/tasks/move",
			body:       `{"action":"move_column","data":{"board_id":3,"column_id":4,"position":2.5}}`,
			wantStatus: http.StatusOK,
			wantCmd:    contract.StageMove{ScopeId: 3, StageId: 4, Position: 2.5},
		},
		{
			name: "move: change_stage", method: http.MethodPost, path: "/tasks/move",
			body:       `{"action":"change_stage","data":{"task_id":5,"board_id":3,"column_id":4,"position":2.5}}`,
			wantStatus: http.StatusOK,
			wantCmd:    contract.WorkItemTransfer{WorkItemId: 5, ScopeId: 3, StageId: 4, Rank: 2.5},
		},
		{
			name: "move: unknown action", method: http.MethodPost, path: "/tasks/move",
			body: `{"action":"foo","data":{}}`, wantStatus: http.StatusBadRequest, wantCode: codeValidation,
		},
		{
			name: "move: missing action", method: http.MethodPost, path: "/tasks/move",
			body: `{"data":{}}`, wantStatus: http.StatusBadRequest, wantCode: codeValidation,
		},
		{
			name: "move: text/plain body", method: http.MethodPost, path: "/tasks/move",
			body:       `{"action":"move_task","data":{"board_id":3,"task_id":5,"position":2.5}}`,
			header:     map[string]string{"Content-Type": "text/plain"},
			wantStatus: http.StatusUnsupportedMediaType, wantCode: codeUnsupportedMediaType,
		},
		{
			name: "move: malformed JSON", method: http.MethodPost, path: "/tasks/move",
			body: `{`, wantStatus: http.StatusBadRequest, wantCode: codeValidation,
		},
		{
			name: "move: data of the wrong shape", method: http.MethodPost, path: "/tasks/move",
			body: `{"action":"move_task","data":"oops"}`, wantStatus: http.StatusBadRequest, wantCode: codeValidation,
		},
		{
			name: "move: missing data", method: http.MethodPost, path: "/tasks/move",
			body: `{"action":"move_task"}`, wantStatus: http.StatusBadRequest, wantCode: codeValidation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app, calls := newTrackerApp(t)
			var opts []test.RequestOption
			for k, v := range tt.header {
				opts = append(opts, test.WithHeader(k, v))
			}
			w, resp := test.Serve(t, newRouter(app), tt.method, tt.path, tt.body, opts...)

			if w.Code != tt.wantStatus {
				t.Errorf("status = %d, want %d\nbody: %s", w.Code, tt.wantStatus, w.Body.String())
			}

			if tt.wantCode == "" {
				if !resp.Status || len(resp.Errors) != 0 {
					t.Errorf("want a successful response, got status=%v errors=%+v", resp.Status, resp.Errors)
				}
			} else {
				if resp.Status {
					t.Errorf("response status = true, want false")
				}
				if len(resp.Errors) == 0 {
					t.Fatalf("want error code %q, got no errors", tt.wantCode)
				}
				if got := resp.Errors[0].Code; got != tt.wantCode {
					t.Errorf("error code = %q, want %q (message: %q)", got, tt.wantCode, resp.Errors[0].Message)
				}
			}

			switch {
			case tt.wantCmd == nil && len(calls.List) != 0:
				t.Errorf("the module must not be reached, got commands %+v", calls.List)
			case tt.wantCmd != nil && len(calls.List) != 1:
				t.Errorf("want exactly one command %+v, got %+v", tt.wantCmd, calls.List)
			case tt.wantCmd != nil && !reflect.DeepEqual(calls.List[0], tt.wantCmd):
				t.Errorf("command = %+v, want %+v", calls.List[0], tt.wantCmd)
			}
		})
	}
}

// Errors that come back from the module keep their own status and code:
// only input the controller rejects itself is a VALIDATION_ERROR.
func TestTaskController_ModuleErrorsKeepTheirStatus(t *testing.T) {
	app := test.NewApp()
	err := bus.RegisterCommand(app.Commands, func(context.Context, contract.WorkItemUpdate) (contract.WorkItemView, error) {
		return contract.WorkItemView{}, errs.NotFound("TASK404", "task not found")
	})
	if err != nil {
		t.Fatal(err)
	}

	w, resp := test.Serve(t, newRouter(app), http.MethodPatch, "/tasks/5", `{"title":"New title"}`)

	if w.Code != http.StatusNotFound {
		t.Errorf("status = %d, want %d", w.Code, http.StatusNotFound)
	}
	if len(resp.Errors) == 0 || resp.Errors[0].Code != "TASK404" {
		t.Errorf("errors = %+v, want code TASK404", resp.Errors)
	}
}
