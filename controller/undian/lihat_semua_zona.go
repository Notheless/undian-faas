package undian

import (
	"main/util"
	"net/http"
)

// LihatSemuaZona function
func (c *Controller) LihatSemuaZona(w http.ResponseWriter, r *http.Request) {
	http := util.NewHandler(w, r)

	res, err := c.srv.LihatSemuaZona(r.Context())

	if err != nil {
		http.ResponseError(err)
		return
	}
	http.ResponseOK(res)
}
