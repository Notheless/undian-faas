// Package p contains an HTTP Cloud Function.
package p

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
)

//EntryPoint it starts here
func EntryPoint(w http.ResponseWriter, r *http.Request) {
	ext := NewHttpx(w, r)
	switch r.Method {
	case http.MethodPost:
		var dok struct {
			Base64   string `json:"file"`
			NamaFile string `json:"namafile"`
		}

		if err := json.NewDecoder(r.Body).Decode(&dok); err != nil {
			ext.ReturnError(err)
			return
		}
		if err := UploadFile(dok.Base64, dok.NamaFile); err != nil {
			ext.ReturnError(err)
			return
		}
		fmt.Fprint(w, html.EscapeString(dok.NamaFile))
	case http.MethodGet:
		db, err := NewDBClient()
		if err != nil {
			ext.ReturnError(err)
			return
		}
		test := r.URL.Query()
		param := test.Get("kategori")
		result, err := GetListPemenang(r.Context(), db, param)
		if err != nil {
			ext.ReturnError(err)
			return
		}
		ext.ReturnJSON(result)
		return

	default:
		fmt.Fprint(w, html.EscapeString(fmt.Sprintf("unsupported method %s ", r.Method)))
	}

	return
}
