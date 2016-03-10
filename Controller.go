package webcontroller

import (
	"github.com/ateleshev/go-webcontext"
)

type Controller struct {
	webcontext.ControllerInterface

	context *webcontext.Context
}

func (this *Controller) Context() *webcontext.Context { // {{{
	return this.context
} // }}}

// [Initialization]

func (this *Controller) Initialize(context *webcontext.Context) { // {{{
	this.context = context
} // }}}

// [Execution]

func (this *Controller) Prepare() error { // {{{
	return nil
} // }}}
