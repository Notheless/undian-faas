package undian

import (
	"context"
	"fmt"
)

func (s *service) LihatPemenangQuery(ctx context.Context, zonaQ []string, kategoriQ []string) ([]PemenangSemuaResult, error) {
	res := []PemenangSemuaResult{}
	var param []interface{}
	var zonaList, katList string
	for _, data := range zonaQ {
		zonaList += "?,"
		param = append(param, data)
	}
	for _, data := range kategoriQ {
		katList += "?,"
		param = append(param, data)

	}

	if len(zonaList) > 0 {
		zonaList = zonaList[:len(zonaList)-1]
	}
	if len(katList) > 0 {
		katList = katList[:len(katList)-1]
	}
	sql := fmt.Sprintf(`SELECT 
	t.nomor
	, t.toko_id 
	, k.nama  
	, k.description 
	, t.zona 
	FROM pemenang p 
	JOIN tiket t ON p.tiket = t.nomor 
	JOIN kategori k ON p.kategori = k.nama  
	WHERE p.deleted = 0
	AND t.zona in (%s)
	AND k.nama in (%s)
	ORDER BY t.zona ASC, k.nama ASC`, zonaList, katList)
	fmt.Println(sql, param)

	rs, err := s.db.QueryContext(ctx, sql, param...)
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
