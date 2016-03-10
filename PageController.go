package webcontroller

import (
	"net/http"
	"os"

	"github.com/ateleshev/go-webcontext"
	"github.com/ateleshev/go-webpage"
)

type PageController struct {
	Controller

	// Page
	Page *webpage.Page

	// Template
	TemplatePath string
	TemplateName string
}

func (this *PageController) SetDefaults() { // {{{
	if this.TemplateName == "" {
		this.TemplateName = webpage.DEFAULT_TEMPLATE_NAME
	}

	if this.TemplatePath == "" {
		this.TemplatePath, _ = os.Getwd()
	}
} // }}}

// [Execution]

func (this *PageController) PageTemplate() *webpage.PageTemplate { // {{{
	this.SetDefaults()
	return webpage.NewPageTemplate(this.TemplateName, this.TemplatePath)
} // }}}

func (this *PageController) Execute(request *http.Request) webcontext.ViewInterface { // {{{
	return NewPageView(this.Context(), request, this.Page, this.PageTemplate())
} // }}}
