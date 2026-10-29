package server

import (
	"net/http"

	"github.com/PolakoJambo/Final-project/pkg/api"
)

func StartServer(port string) error {
	api.Init()

	http.Handle("/", http.FileServer(http.Dir("web")))

	return http.ListenAndServe(":"+port, nil)
}
