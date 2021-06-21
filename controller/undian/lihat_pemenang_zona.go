package undian

import (
	"main/util"
	"net/http"
)

// LihatPemenangZonasi function
func (c *Controller) LihatPemenangZonasi(w http.ResponseWriter, r *http.Request) {
	http := util.NewHandler(w, r)

	zona := http.RouteValue("zona")
	res, err := c.srv.LihatPemenangZonasi(r.Context(), zona)

	if err != nil {
		http.ResponseError(err)
		return
	}
	http.ResponseOK(res)

}
