package utils

import (
	"bytes"
	"context"

	"github.com/a-h/templ"
)

func TemplToString(ctx context.Context, tmpl templ.Component) string {
	var buf bytes.Buffer
	err := tmpl.Render(ctx, &buf)
	if err != nil {
		panic(err)
	}
	return buf.String()
}
