package undian

import (
	"net/http"
)

// GetUndian function
func (c *Controller) GetUndian(w http.ResponseWriter, r *http.Request) {

	err := c.srv.GetUndian(r.Context())

	if err != nil {

		return
	}

}
