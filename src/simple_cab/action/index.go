package action

import (
	"fmt"
	"net/http"
)

func Index(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "hello SimpleCab")
}
