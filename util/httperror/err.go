package httperror

import (
	"net/http"
	"strings"
)

func Message(w http.ResponseWriter, r *http.Request, msg string) {
	url := []string{"/err?msg=", msg}
	http.Redirect(w, r, strings.Join(url, ""), http.StatusFound)
}
