package undian

import (
	"main/util"
	"net/http"
)

// LihatSemuaPemenang function
func (c *Controller) LihatSemuaPemenang(w http.ResponseWriter, r *http.Request) {
	http := util.NewHandler(w, r)

	res, err := c.srv.LihatSemuaPemenang(r.Context())

	if err != nil {
		http.ResponseError(err)
		return
	}
	http.ResponseOK(res)

}
