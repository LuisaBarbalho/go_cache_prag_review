package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

type chiRouter struct{}

var (
	chiDispatcher = chi.NewRouter()
)

func NewChiRouter() Router {
	return &chiRouter{}
}

func (*chiRouter) GET(uri string, f func(response http.ResponseWriter, request *http.Request)) {
	chiDispatcher.Get(uri, f)
}
func (*chiRouter) POST(uri string, f func(response http.ResponseWriter, request *http.Request)) {
	chiDispatcher.Post(uri, f)
}
func (*chiRouter) SERVE(port string) {
	fmt.Printf("Chi HTTP server running on port %v", port)
	http.ListenAndServe(port, chiDispatcher)
}
