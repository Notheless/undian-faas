// Package p contains an HTTP Cloud Function.
package p

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//EntryPoint it starts here
func EntryPoint(w http.ResponseWriter, r *http.Request) {
	ext := NewHttpx(w, r)
	switch r.Method {
	case http.MethodPost:
		db, err := NewDBClient()
		if err != nil {
			ext.ReturnError(err)
			return
		}
		dok := &Dokumen{}
		if err := json.NewDecoder(r.Body).Decode(&dok); err != nil {
			ext.ReturnError(err)
			return
		}
		if err := UploadFile(dok.Base64, dok.NamaFile); err != nil {
			ext.ReturnError(err)
			return
		}
		if err := ProcessExcel(r.Context(), db, dok.Base64); err != nil {
			ext.ReturnError(err)
			return
		}
		if err := GeneratePemenang(r.Context(), db); err != nil {
			ext.ReturnError(err)
			return
		}
		ext.ReturnText("OK")
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
		ext.ReturnError(fmt.Errorf("Method %s is implemented yet", r.Method))
	}

	return
}

//Dokumen struk
type Dokumen struct {
	Base64   string `json:"file"`
	NamaFile string `json:"namafile"`
}
