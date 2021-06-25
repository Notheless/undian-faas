package undian

import (
	"main/util"
	"net/http"
)

// CleanPemenang function
func (c *Controller) CleanPemenang(w http.ResponseWriter, r *http.Request) {
	http := util.NewHandler(w, r)
	err := http.VerifyKey()
	if err != nil {
		http.ResponseError(err)
		return
	}

	err = c.srv.CleanPemenang(r.Context())

	if err != nil {
		http.ResponseError(err)
		return
	}
	http.ResponseOK("OK")

}
