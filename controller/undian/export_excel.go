package undian

import (
	"fmt"
	"main/util"
	"net/http"
	"time"
)

// ExportExcel function
func (c *Controller) ExportExcel(w http.ResponseWriter, r *http.Request) {
	http := util.NewHandler(w, r)

	res, err := c.srv.ExportExcel(r.Context())

	if err != nil {
		http.ResponseError(err)
		return
	}
	fileName := "\"List-Pemenang-Lottery " + time.Now().Local().Format("2006-01-02") + ".xlsx\""
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment;filename=%s", fileName))
	http.ResponseOK(res)

}
