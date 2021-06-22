package undian

import (
	"main/util"
	"net/http"
)

// GeneratePemenangQuery function
func (c *Controller) GeneratePemenangQuery(w http.ResponseWriter, r *http.Request) {
	http := util.NewHandler(w, r)
	err := http.VerifyKey()
	if err != nil {
		http.ResponseError(err)
		return
	}

	zona := http.QueryValue("zona")
	kategori := http.QueryValue("kategori")
	res, err := c.srv.GeneratePemenangQuery(r.Context(), zona, kategori)

	if err != nil {
		http.ResponseError(err)
		return
	}
	http.ResponseOK(res)

}
