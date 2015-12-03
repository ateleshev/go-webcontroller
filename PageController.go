package webcontroller

import (
	"net/http"
	"os"

	"github.com/ArtemTeleshev/go-webcontext"
	"github.com/ArtemTeleshev/go-webpage"
)

type PageController struct {
	Controller

	// Page
	Page *webpage.Page

	// Template
	TemplatePath string
	TemplateName string
}

// [Execution]

func (this *PageController) PageTemplate() *webpage.PageTemplate { // {{{
	if this.TemplateName == "" {
		this.TemplateName = webpage.DEFAULT_TEMPLATE_NAME
	}

	if this.TemplatePath == "" {
		this.TemplatePath, _ = os.Getwd()
	}

	return webpage.NewPageTemplate(this.TemplateName, this.TemplatePath)
} // }}}

func (this *PageController) Execute(request *http.Request) webcontext.ViewInterface { // {{{
	return NewPageView(this.Context(), request, this.Page, this.PageTemplate())
} // }}}
