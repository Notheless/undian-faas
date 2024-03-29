package undian

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func (s *service) GeneratePemenangQuery(ctx context.Context, zonaQ []string, kategoriQ []string) ([]PemenangSemuaResult, error) {
	for _, zona := range zonaQ {
		for _, kategori := range kategoriQ {

			//cek zona
			zonaExist := false
			cekZonaSQL := "SELECT CASE WHEN count(z.id) < 1 THEN FALSE ELSE TRUE END FROM zona z WHERE z.id = ?"
			err := s.db.QueryRowContext(ctx, cekZonaSQL, zona).Scan(&zonaExist)
			if err != nil {
				return nil, err
			}
			if !zonaExist {
				return nil, fmt.Errorf("Zona %s tidak exist", zona)
			}

			//cek kategori
			maksimalPemenang := 0
			cekKategoriSQL := "SELECT k.jumlah_pemenang FROM kategori k WHERE k.nama = ?"
			err = s.db.QueryRowContext(ctx, cekKategoriSQL, kategori).Scan(&maksimalPemenang)
			if err != nil {
				return nil, err
			}
			if maksimalPemenang == 0 {
				return nil, fmt.Errorf("kategori %s tidak ada", kategori)
			}

			//get list peserta
			listNomorUndian := []string{}
			ambilNomorSQL := "SELECT t.nomor FROM tiket t WHERE t.zona  = ? AND t.nomor NOT IN (SELECT p.tiket FROM pemenang p WHERE p.deleted = 0 AND p.zona = ? AND p.kategori != ?)"
			rs, err := s.db.QueryContext(ctx, ambilNomorSQL, zona, zona, kategori)
			if err != nil {
				return nil, err
			}
			for rs.Next() {
				var str string
				rs.Scan(&str)
				listNomorUndian = append(listNomorUndian, str)
			}

			//proses random pemenang
			listVal := ""
			var listArg []interface{}
			for i := 0; i < maksimalPemenang && len(listNomorUndian) != 0; i++ {
				j := int64(rand.Intn(len(listNomorUndian))) * int64(rand.Intn(time.Now().Nanosecond())) % int64(len(listNomorUndian))
				listArg = append(listArg, listNomorUndian[j])
				listArg = append(listArg, kategori)
				listArg = append(listArg, zona)
				listNomorUndian = remove(listNomorUndian, int(j))
				listVal += "(?, ?, ?),"
			}
			listVal = listVal[:len(listVal)-1]
			InsertSQL := fmt.Sprintf("INSERT INTO pemenang (tiket, kategori, zona) VALUES %s;", listVal)

			//delete pemenang lama
			DeleteSQL := "UPDATE pemenang SET deleted = true WHERE kategori = ? AND zona = ?"
			_, err = s.db.ExecContext(ctx, DeleteSQL, kategori, zona)
			if err != nil {
				return nil, err
			}
			//insert pemenang baru
			_, err = s.db.ExecContext(ctx, InsertSQL, listArg...)
			if err != nil {
				return nil, err
			}

		}
	}
	//Get pemenang baru
	result, err := s.LihatPemenangQuery(ctx, zonaQ, kategoriQ)

	//return
	return result, err
}
