package app

import (
	"fmt"
	"github.com/teukumulya-ichsan/couchDB-go/router"
	"net/http"
)

var (
	httpRouter = router.NewChiRouter()
)

func StartApplication() {
	const port = ":3000"

	httpRouter.GET("/", func(res http.ResponseWriter, req *http.Request) {
		_, _ = fmt.Fprintln(res, "Up and Running...")
	})

	mapsUrls()
	httpRouter.SERVE(port)
}
