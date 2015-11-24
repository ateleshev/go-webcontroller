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

func (this *PageController) Initialize(context *webcontext.Context) { // {{{
	this.Controller.Initialize(context)

	// Template

	if len(this.TemplateName) == 0 {
		this.TemplateName = webpage.DEFAULT_TEMPLATE_NAME
	}

	if len(this.TemplatePath) == 0 {
		this.TemplatePath, _ = os.Getwd()
	}
} // }}}

func (this *PageController) Render(writer http.ResponseWriter) error { // {{{
	template := webpage.NewPageTemplate(this.TemplateName, this.TemplatePath)
	if err := template.Execute(writer, this.Page); err != nil {
		return err
	}

	return nil
} // }}}
