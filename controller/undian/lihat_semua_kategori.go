package undian

import (
	"main/util"
	"net/http"
)

// LihatSemuaKategori function
func (c *Controller) LihatSemuaKategori(w http.ResponseWriter, r *http.Request) {
	http := util.NewHandler(w, r)

	res, err := c.srv.LihatSemuaKategori(r.Context())

	if err != nil {
		http.ResponseError(err)
		return
	}
	http.ResponseOK(res)
}
