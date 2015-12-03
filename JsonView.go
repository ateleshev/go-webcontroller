package webcontroller

import (
	"encoding/json"
	"net/http"

	"github.com/ArtemTeleshev/go-webcontext"
)

const (
	CONTENT_TYPE_JSON = "application/json"
)

func NewJsonView(context *webcontext.Context, request *http.Request, data interface{}) *JsonView { // {{{
	result := &JsonView{data: data}
	result.Initialize(context, request)

	return result
} // }}}

type JsonView struct {
	View

	data interface{}
}

func (this *JsonView) Render(responseWriter http.ResponseWriter) error { // {{{
	responseWriter.Header().Set("Content-Type", CONTENT_TYPE_JSON)
	// Write headers
	responseWriter.WriteHeader(http.StatusOK)

	json, err := json.Marshal(this.data)
	if err != nil {
		return err
	}

	// Write body
	responseWriter.Write(json)

	return nil
} // }}}
