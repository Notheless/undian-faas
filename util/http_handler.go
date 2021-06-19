package util

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	contentTypeBinary = "application/octet-stream"
	contentTypeForm   = "application/x-www-form-urlencoded"
	contentTypeJSON   = "application/json"
	contentTypeHTML   = "text/html; charset=utf-8"
	contentTypeText   = "text/plain; charset=utf-8"
)

//HTTPHandler struct
type HTTPHandler struct {
	w    http.ResponseWriter
	r    *http.Request
	vars map[string]string
}

// NewHandler func
func NewHandler(w http.ResponseWriter, r *http.Request) *HTTPHandler {
	vars := mux.Vars(r)
	return &HTTPHandler{w: w, r: r, vars: vars}
}

//ResponseOK function
func (h *HTTPHandler) ResponseOK(body interface{}) {
	payload, _ := json.Marshal(body)
	h.w.Header().Set("Content-Type", contentTypeJSON)
	if payload[0] == '"' {
		fmt.Println(string(payload))
		payload = []byte(fmt.Sprintf("%v", body))
		h.w.Header().Set("Content-Type", contentTypeText)
	}
	h.w.WriteHeader(http.StatusOK)
	h.w.Write(payload)
}

//ResponseError function
func (h *HTTPHandler) ResponseError(body error) {
	h.w.Header().Add("Content-Type", contentTypeText)
	h.w.WriteHeader(http.StatusBadRequest)
	h.w.Write([]byte(body.Error()))
}

//RouteValue function
func (h *HTTPHandler) RouteValue(key string) string {
	return h.vars[key]
}

//QueryValue function
func (h *HTTPHandler) QueryValue(key string) string {
	return h.r.URL.Query().Get(key)
}

//BodyParse function
func (h *HTTPHandler) BodyParse(b interface{}) error {
	decoder := json.NewDecoder(h.r.Body)
	return decoder.Decode(b)
}
