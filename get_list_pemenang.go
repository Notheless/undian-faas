package p

import (
	"context"
	"database/sql"
)

//GetListPemenang function
func GetListPemenang(ctx context.Context, db *sql.DB, kategori string) (interface{}, error) {
	sql := `SELECT 
		ku.nama_kategori 
		,dn.nomor_undian 
		FROM daftar_pemenang dp 
		LEFT JOIN daftar_nomor dn ON dp.nomor_undian = dn.id 
		LEFT JOIN kategori_undian ku ON dp.kategori = ku.id`
	if kategori != "" {
		result := DaftarPemenang{Kategori: kategori, Pemenang: []string{}}
		sql += `
		WHERE ku.nama_kategori = ?`
		rs, err := db.QueryContext(ctx, sql, kategori)
		if err != nil {
			return nil, err
		}
		for rs.Next() {
			var namakat string
			var nomor string
			rs.Scan(&namakat, &nomor)
			result.Pemenang = append(result.Pemenang, nomor)
		}
		return result, nil
	}
	result := []DaftarPemenang{}
	rs, err := db.QueryContext(ctx, sql, kategori)
	if err != nil {
		return nil, err
	}
	for rs.Next() {
		var namakat string
		var nomor string
		rs.Scan(&namakat, &nomor)
		found := false
		for i, data := range result {
			if data.Kategori == namakat {
				result[i].Pemenang = append(result[i].Pemenang, nomor)
				found = true
			}
		}
		if !found {
			newKat := DaftarPemenang{
				Kategori: namakat,
				Pemenang: []string{nomor},
			}
			result = append(result, newKat)
		}
	}
	return result, nil
}

//DaftarPemenang struct
type DaftarPemenang struct {
	Kategori string   `json:"kategori"`
	Pemenang []string `json:"pemenang"`
}
