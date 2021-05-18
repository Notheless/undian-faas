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

//EntryPoint it starts here
func EntryPoint(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
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
	case http.MethodGet:
		db, err := NewDBClient()
		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusBadRequest)
			return
		}
		test := r.URL.Query()
		param := test.Get("kategori")
		result, err := GetListPemenang(r.Context(), db, param)
		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusBadRequest)
			return
		}
		fmt.Fprint(w, html.EscapeString(ConvertJSON(result)))

	default:
		fmt.Fprint(w, html.EscapeString(fmt.Sprintf("unsupported method %s ", r.Method)))
	}

	return
}

//ConvertJSON function
func ConvertJSON(in interface{}) string {
	data, _ := json.Marshal(in)
	return string(data)
}
