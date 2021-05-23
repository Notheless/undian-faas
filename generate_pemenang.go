package p

import (
	"context"
	"database/sql"
	"fmt"
	"math/rand"
)

//GeneratePemenang function
func GeneratePemenang(ctx context.Context, db *sql.DB) error {
	logger := CreateLogger(ctx)
	listKategori := []Kategori{}
	listNomorUndian := []NomorUndian{}
	var listArg []interface{}
	SQLgetKategori := "SELECT id, nama_kategori, maksimal_pemenang FROM kategori_undian;"
	SQLgetNomor := "SELECT id, nomor_undian FROM `default`.daftar_nomor;"
	logger.Println("Get list Kategori")
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

	logger.Println("Get list nomor")
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
		if len(listNomorUndian) == 0 {
			break
		}
		for i := 1; i < katergori.Maksimal || len(listNomorUndian) == 0; i++ {
			j := rand.Intn(len(listNomorUndian))
			listArg = append(listArg, listNomorUndian[j].ID)
			listArg = append(listArg, katergori.ID)
			listNomorUndian = remove(listNomorUndian, j)
			listVal += "('?', '?'),"
		}
	}

	if len(listVal) == 0 {
		return fmt.Errorf("Tidak ada data")
	}
	listVal = listVal[:len(listVal)-1]
	SQLInsert := fmt.Sprintf("INSERT INTO daftar_pemenang (nomor_undian, kategori) VALUES %s;", listVal)

	logger.Println("Insert Pemenang")
	_, err = db.QueryContext(ctx, SQLInsert, listArg...)
	if err != nil {
		return err
	}

	return nil
}
func remove(s []NomorUndian, i int) []NomorUndian {
	s[i] = s[len(s)-1]
	// We do not need to put s[i] at the end, as it will be discarded anyway
	return s[:len(s)-1]
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
