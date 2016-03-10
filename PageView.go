package webcontroller

import (
	"net/http"

	"github.com/ateleshev/go-webcontext"
	"github.com/ateleshev/go-webpage"
)

func NewPageView(context *webcontext.Context, request *http.Request, page *webpage.Page, template *webpage.PageTemplate) *PageView { // {{{
	result := &PageView{
		page:     page,
		template: template,
	}
	result.Initialize(context, request)

	return result
} // }}}

type PageView struct {
	View

	page     *webpage.Page
	template *webpage.PageTemplate
}

// [Render]

func (this *PageView) Render(responseWriter http.ResponseWriter) error { // {{{
	if err := this.template.Execute(responseWriter, this.page); err != nil {
		return err
	}

	return nil
} // }}}
