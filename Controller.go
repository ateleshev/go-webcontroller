package webcontroller

import (
	"net/http"

	"github.com/ArtemTeleshev/go-webcontext"
)

type Controller struct {
	webcontext.ControllerInterface

	// [Protected]
	context *webcontext.Context
	request *http.Request

	// [Public]
	Err error
}

// [Error]

func (this *Controller) HasError() bool { // {{{
	return this.Err != nil
} // }}}

func (this *Controller) Error() string { // {{{
	return this.Err.Error()
} // }}}

// [Context]

func (this *Controller) HasContext() bool { // {{{
	return this.context != nil
} // }}}

func (this *Controller) Context() *webcontext.Context { // {{{
	return this.context
} // }}}

// [Request]

func (this *Controller) HasRequest() bool { // {{{
	return this.request != nil
} // }}}

func (this *Controller) Request() *http.Request { // {{{
	return this.request
} // }}}

// [Controller]

func (this *Controller) Initialize(context *webcontext.Context) { // {{{
	this.context = context
} // }}}

func (this *Controller) Configure(request *http.Request) { // {{{
	this.request = request
} // }}}

func (this *Controller) Prepare() error { // {{{
	return nil
} // }}}

// [Error]

func (this *Controller) RenderError(writer http.ResponseWriter) { // {{{
	http.Error(writer, "Forbidden", http.StatusForbidden)
} // }}}
