package app

import (
	"fmt"
	"net/http"

	"github.com/teukumulya-ichsan/couchDB-go/router"
)

var (
	httpRouter = router.NewChiRouter()
)

// StartApplication function to started serve http ...
func StartApplication() {
	const port = ":3000"

	httpRouter.GET("/", func(res http.ResponseWriter, req *http.Request) {
		_, _ = fmt.Fprintln(res, "Up and Running...")
	})

	mapsUrls()
	httpRouter.SERVE(port)
}
