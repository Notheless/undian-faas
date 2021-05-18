package p

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	ctjson = "application/json; charset=UTF-8"
	cttext = "text/plain; charset=utf-8"
)

//HTTPx struct
type HTTPx struct {
	wr http.ResponseWriter
	rq *http.Request
}

//NewHttpx func
func NewHttpx(w http.ResponseWriter, r *http.Request) *HTTPx {
	return &HTTPx{wr: w, rq: r}
}

//ConvertJSON function
func ConvertJSON(in interface{}) string {
	data, _ := json.Marshal(in)
	return string(data)
}

//ReturnJSON function
func (h *HTTPx) ReturnJSON(val interface{}) {
	h.ResponseReturn(http.StatusOK, ctjson, ConvertJSON(val))
}

//ReturnText function
func (h *HTTPx) ReturnText(val string) {
	h.ResponseReturn(http.StatusOK, cttext, val)
}

//ReturnError function
func (h *HTTPx) ReturnError(val error) {
	h.ResponseReturn(http.StatusBadRequest, cttext, fmt.Sprint(val))
}

//ResponseReturn func
func (h *HTTPx) ResponseReturn(code int, contenttype string, value string) {
	h.wr.Header().Set("Content-Type", contenttype)
	h.wr.Write([]byte(value))
	h.wr.WriteHeader(code)
}
