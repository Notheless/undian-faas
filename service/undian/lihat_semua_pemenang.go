package undian

import (
	"context"
)

func (s *service) LihatSemuaPemenang(ctx context.Context) ([]PemenangSemuaResult, error) {
	res := []PemenangSemuaResult{}
	sql := `SELECT 
	t.nomor
	, t2.nama 
	, k.nama  
	, k.description 
	, t.zona 
	FROM pemenang p 
	JOIN tiket t ON p.tiket = t.nomor 
	JOIN kategori k ON p.kategori = k.nama  
	JOIN toko t2 ON t.toko_id = t2.customer_id 
	WHERE p.deleted = 0
	ORDER BY t.zona ASC, p.kategori ASC`

	rs, err := s.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	for rs.Next() {
		pem := &PemenangResult{}
		var kategori, desc, zona string
		rs.Scan(
			&pem.Tiket,
			&pem.NamaToko,
			&kategori,
			&desc,
			&zona,
		)
		newZona := true
		for i, zonaData := range res {
			if zona == zonaData.Zona {
				newKategori := true
				for j, data := range res[i].ZonaPemenang {
					if kategori == data.Kategori {
						res[i].ZonaPemenang[j].Pemenang = append(res[i].ZonaPemenang[j].Pemenang, *pem)
						newKategori = false
						break
					}
				}
				if newKategori {
					pzr := PemenangZonaResult{
						Pemenang:            []PemenangResult{*pem},
						Kategori:            kategori,
						KategoriDescription: desc,
					}
					res[i].ZonaPemenang = append(res[i].ZonaPemenang, pzr)
				}
				newZona = false
				break
			}
		}
		if newZona {

			pzr := PemenangZonaResult{
				Pemenang:            []PemenangResult{*pem},
				Kategori:            kategori,
				KategoriDescription: desc,
			}
			zon := PemenangSemuaResult{ZonaPemenang: []PemenangZonaResult{pzr}, Zona: zona}
			res = append(res, zon)
		}

	}
	return res, nil
}
