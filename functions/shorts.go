package shorts

import (
	"fmt"
	"net/http"
)

func NotExist(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "{\"e\":1,\"d\":[1,404]}")
}
