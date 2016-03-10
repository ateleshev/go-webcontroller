package webcontroller

import (
	"github.com/ateleshev/go-webcontext"
	"net/http"
)

type View struct {
	webcontext.ViewInterface

	// == Protected ===
	context *webcontext.Context
	request *http.Request
}

func (this *View) Context() *webcontext.Context { // {{{
	return this.context
} // }}}

func (this *View) Request() *http.Request { // {{{
	return this.request
} // }}}

// [Initialization]

func (this *View) Initialize(context *webcontext.Context, request *http.Request) { // {{{
	this.context = context
	this.request = request
} // }}}

// [Render]

func (this *View) Render(responseWriter http.ResponseWriter) error { // {{{
	http.Error(responseWriter, "Forbidden", http.StatusForbidden)

	return nil
} // }}}
