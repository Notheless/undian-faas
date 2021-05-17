// Package p contains an HTTP Cloud Function.
package p

import (
	"encoding/json"
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
)

// HelloWorld prints the JSON encoded "message" field in the body
// of the request or "Hello, World!" if there isn't one.
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	var dok struct {
		Base64   string `json:"file"`
		NamaFile string `json:"namafile"`
	}

	if err := json.NewDecoder(r.Body).Decode(&dok); err != nil {
		switch err {
		case io.EOF:
			http.Error(w, "end of file error", http.StatusBadRequest)
			return
		default:
			log.Printf("json.NewDecoder: %v", err)
			http.Error(w, fmt.Sprint(err), http.StatusBadRequest)
			return
		}
	}
	if err := UploadFile(dok.Base64, dok.NamaFile); err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusBadRequest)
		return
	}
	fmt.Fprint(w, html.EscapeString(dok.NamaFile))
	return
}
