package undian

import (
	"context"
)

func (s *service) LihatPemenangZonasi(ctx context.Context, zona string) ([]PemenangZonaResult, error) {
	res := []PemenangZonaResult{}
	sql := `SELECT 
	t.nomor
	, t.toko_id 
	, k.nama  
	, k.description 
	FROM pemenang p 
	JOIN tiket t ON p.tiket = t.nomor 
	JOIN kategori k ON p.kategori = k.nama 
	WHERE p.zona = ? AND p.deleted = 0`

	rs, err := s.db.QueryContext(ctx, sql, zona)
	if err != nil {
		return nil, err
	}
	for rs.Next() {
		pem := &PemenangResult{}
		var kategori, desc string
		rs.Scan(
			&pem.Tiket,
			&pem.NamaToko,
			&kategori,
			&desc,
		)
		newData := true
		for i, data := range res {
			if kategori == data.Kategori {
				res[i].Pemenang = append(res[i].Pemenang, *pem)
				newData = false
				break
			}
		}
		if newData {
			pzr := PemenangZonaResult{
				Pemenang:            []PemenangResult{*pem},
				Kategori:            kategori,
				KategoriDescription: desc,
			}
			res = append(res, pzr)
		}
	}
	return res, nil
}
