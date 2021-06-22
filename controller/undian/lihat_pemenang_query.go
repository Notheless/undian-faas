package undian

import (
	"main/util"
	"net/http"
)

// LihatPemenangQuery function
func (c *Controller) LihatPemenangQuery(w http.ResponseWriter, r *http.Request) {
	http := util.NewHandler(w, r)
	zona := http.QueryValue("zona")
	kategori := http.QueryValue("kategori")
	res, err := c.srv.LihatPemenangQuery(r.Context(), zona, kategori)

	if err != nil {
		http.ResponseError(err)
		return
	}
	http.ResponseOK(res)

}
