package webcontroller

import (
	"net/http"
	"path"

	"github.com/gorilla/mux" // http://www.gorillatoolkit.org/pkg/mux

	"github.com/ArtemTeleshev/go-webcontext"
	"github.com/ArtemTeleshev/go-webpage"
)

type WebContentController struct {
	PageController
}

// [Initialization]

func (this *WebContentController) Register(route *mux.Route) { // {{{
	route.PathPrefix("/")
} // }}}

// [Execution]

func (this *WebContentController) Execute(request *http.Request) webcontext.ViewInterface { // {{{
	this.SetDefaults()
	dir := path.Join(this.TemplatePath, webpage.TEMPLATE_DIR_MAIN, this.TemplateName, webpage.TEMPLATE_DIR_WEB)
	return NewWebContentView(this.Context(), request, dir)
} // }}}
