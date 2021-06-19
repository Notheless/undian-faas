package undian

import (
	"main/util"
	"net/http"
)

// GeneratePemenang function
func (c *Controller) GeneratePemenang(w http.ResponseWriter, r *http.Request) {
	http := util.NewHandler(w, r)

	zona := http.RouteValue("zona")
	kategori := http.RouteValue("kategori")
	res, err := c.srv.GeneratePemenang(r.Context(), zona, kategori)

	if err != nil {
		http.ResponseError(err)
		return
	}
	http.ResponseOK(res)

}
