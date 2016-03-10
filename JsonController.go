package webcontroller

import (
	"net/http"

	"github.com/ateleshev/go-webcontext"
)

type JsonController struct {
	Controller

	Data interface{}
}

func (this *JsonController) Execute(request *http.Request) webcontext.ViewInterface { // {{{
	return NewJsonView(this.Context(), request, this.Data)
} // }}}
