package undian

import (
	"context"
)

func (s *service) LihatSemuaKategori(ctx context.Context) ([]KategoriResult, error) {
	var res []KategoriResult

	sql := `SELECT nama,description,jumlah_pemenang FROM undian.kategori`
	rs, err := s.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	for rs.Next() {
		kat := &KategoriResult{}
		rs.Scan(
			&kat.Nama,
			&kat.Description,
			&kat.Jumlah)
		res = append(res, *kat)
	}
	return res, nil
}
