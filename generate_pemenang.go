package p

import (
	"context"
	"database/sql"
	"fmt"
)

//GeneratePemenang function
func GeneratePemenang(ctx context.Context, db *sql.DB) error {
	listKategori := []Kategori{}
	listNomorUndian := []NomorUndian{}
	var listArg []interface{}
	SQLgetKategori := "SELECT id, nama_kategori, maksimal_pemenang FROM kategori_undian;"
	SQLgetNomor := "SELECT id, nomor_undian FROM `default`.daftar_nomor;"
	rs, err := db.QueryContext(ctx, SQLgetKategori)
	if err != nil {
		return err
	}
	for rs.Next() {
		res := Kategori{}
		rs.Scan(
			&res.ID,
			&res.Nama,
			&res.Maksimal,
		)
		listKategori = append(listKategori, res)
	}

	rs, err = db.QueryContext(ctx, SQLgetNomor)
	if err != nil {
		return err
	}
	for rs.Next() {
		res := NomorUndian{}
		rs.Scan(
			&res.ID,
			&res.Nomor,
		)
		listNomorUndian = append(listNomorUndian, res)
	}
	listVal := ""

	for _, katergori := range listKategori {
		for i := 1; i < katergori.Maksimal; i++ {
			j := 0
			listArg = append(listArg, listNomorUndian[j].ID)
			listArg = append(listArg, katergori.ID)
			listVal += "('?', '?'),"
		}
	}

	if len(listVal) == 0 {
		return fmt.Errorf("Tidak ada data")
	}
	listVal = listVal[:len(listVal)-1]
	SQLInsert := fmt.Sprintf("INSERT INTO daftar_pemenang (nomor_undian, kategori) VALUES %s;", listVal)

	_, err = db.QueryContext(ctx, SQLInsert, listArg...)
	if err != nil {
		return err
	}

	return nil
}

type (
	Kategori struct {
		ID       int
		Nama     string
		Maksimal int
	}
	NomorUndian struct {
		ID    int
		Nomor string
	}
	Pemenang struct {
		IDkategori int
		IDNomor    int
	}
)
