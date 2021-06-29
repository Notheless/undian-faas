package undian

import (
	"context"
)

func (s *service) LihatPemenang(ctx context.Context, zona string, kategori string) ([]PemenangResult, error) {
	res := []PemenangResult{}
	sql := `SELECT 
	t.nomor
	, t2.nama 
	FROM pemenang p 
	JOIN tiket t ON p.tiket = t.nomor 
	JOIN toko t2 ON t.toko_id = t2.customer_id 
	WHERE p.zona = ? AND p.kategori = ? AND p.deleted = 0`

	rs, err := s.db.QueryContext(ctx, sql, zona, kategori)
	if err != nil {
		return nil, err
	}
	for rs.Next() {
		pem := &PemenangResult{}
		rs.Scan(
			&pem.Tiket,
			&pem.NamaToko,
		)
		res = append(res, *pem)
	}
	return res, nil
}
