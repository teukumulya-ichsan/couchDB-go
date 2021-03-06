package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

type chiRouter struct {
	chi.Router
}

// NewChiRouter constructor ...
func NewChiRouter() Router {
	chiDispatcher := chi.NewRouter()
	return &chiRouter{chiDispatcher}
}

func (c chiRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	c.Get(uri, f)
}

func (c chiRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	c.Post(uri, f)
}

func (c chiRouter) SERVE(port string) {
	fmt.Printf("Chi HTPP Server running on port : %v", port)
	_ = http.ListenAndServe(port, c)
}
