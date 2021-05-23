// Package p contains an HTTP Cloud Function.
package p

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"cloud.google.com/go/logging"
)

//EntryPoint it starts here
func EntryPoint(w http.ResponseWriter, r *http.Request) {

	// Sets your Google Cloud Platform project ID.
	projectID := "bold-camera-314007"
	ctx := r.Context()
	// Creates a client.
	client, err := logging.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()
	// Sets the name of the log to write to.
	logName := "my-log"

	logger := client.Logger(logName).StandardLogger(logging.Debug)

	// Logs "hello world", log entry is visible at
	// Cloud Logs.
	ext := NewHttpx(w, r)
	switch r.Method {
	case http.MethodPost:
		db, err := NewDBClient()
		if err != nil {
			ext.ReturnError(err)
			return
		}
		dok := &Dokumen{}
		logger.Println("Decode")
		if err := json.NewDecoder(r.Body).Decode(&dok); err != nil {
			ext.ReturnError(err)
			return
		}
		logger.Println("Upload")
		if err := UploadFile(dok.Base64, dok.NamaFile); err != nil {
			ext.ReturnError(err)
			return
		}
		logger.Println("Excel parse")
		if err := ProcessExcel(ctx, db, dok.Base64); err != nil {
			ext.ReturnError(err)
			return
		}
		logger.Println("Update")
		if err := GeneratePemenang(ctx, db); err != nil {
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
		result, err := GetListPemenang(ctx, db, param)
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
