package webcontroller

import (
	"net/http"

	"github.com/ArtemTeleshev/go-webcontext"
)

func NewWebContentView(context *webcontext.Context, request *http.Request, dir string) *WebContentView { // {{{
	result := &WebContentView{dir: dir}
	result.Initialize(context, request)

	return result
} // }}}

type WebContentView struct {
	View

	dir string
}

// [Render]

func (this *WebContentView) Render(responseWriter http.ResponseWriter) error { // {{{
	http.FileServer(http.Dir(this.dir)).ServeHTTP(responseWriter, this.Request())

	return nil
} // }}}
