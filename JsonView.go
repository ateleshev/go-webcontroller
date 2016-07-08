package webcontroller

import (
	"compress/gzip"
	"net/http"
	"strings"

	encoder "encoding/json"

	"github.com/ateleshev/go-webcontext"
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

	acceptEncoding := this.Request().Header.Get("Accept-Encoding")
	if strings.Contains(acceptEncoding, CONTEN_ENCODING_GZIP) {
		// responseWriter.Header().Add("Content-Type", CONTENT_TYPE_XGZIP)
		responseWriter.Header().Set("Content-Encoding", CONTEN_ENCODING_GZIP)
	}

	// Write headers
	responseWriter.WriteHeader(http.StatusOK)

	json, err := encoder.Marshal(this.data)
	if err != nil {
		return err
	}

	// Write body
	if strings.Contains(acceptEncoding, CONTEN_ENCODING_GZIP) {
		gz := gzip.NewWriter(responseWriter)
		defer gz.Close()
		gz.Write(json)
	} else {
		responseWriter.Write(json)
	}

	return nil
} // }}}
