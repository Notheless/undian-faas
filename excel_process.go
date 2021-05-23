package p

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/base64"
	"fmt"
	"log"

	"github.com/360EntSecGroup-Skylar/excelize"
)

//ProcessExcel function
func ProcessExcel(ctx context.Context, db *sql.DB, logger *log.Logger, FileBase64 string) error {
	dec, err := base64.StdEncoding.DecodeString(FileBase64)
	if err != nil {
		return err
	}
	sqlDeleteNomor := "DELETE FROM daftar_nomor;"
	sqlDeleteDaftar := "DELETE FROM daftar_pemenang;"
	listVal := ""
	f, _ := excelize.OpenReader(bytes.NewReader(dec))
	endOfFile := false
	var listNomorUndian []interface{}
	logger.Println("Proses data")
	for i := 2; !endOfFile; i++ {
		cell := f.GetCellValue("Sheet1", "A"+fmt.Sprintf("%d", i))
		listNomorUndian = append(listNomorUndian, cell)
		listVal += "'?',"
		logger.Printf("Reading data no : %d \n\r", i)
	}
	if len(listVal) == 0 {
		return fmt.Errorf("Tidak ada data")
	}
	logger.Println("Selesai Proses data")
	listVal = listVal[:len(listVal)-1]

	sqlInsert := fmt.Sprintf(`INSERT INTO daftar_nomor (nomor_undian) VALUES (%s)`, listVal)
	logger.Println("Delete Nomor")
	_, err = db.ExecContext(ctx, sqlDeleteNomor)
	if err != nil {
		return err
	}
	logger.Println("Delete Pemenang")
	_, err = db.ExecContext(ctx, sqlDeleteDaftar)
	if err != nil {
		return err
	}
	logger.Println("Insert Nomor")
	_, err = db.ExecContext(ctx, sqlInsert, listNomorUndian...)
	if err != nil {
		return err
	}
	logger.Println("Excel selesai")
	return nil
}
