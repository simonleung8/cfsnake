package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	PortVar = "VCAP_APP_PORT"
)

type player struct {
	dice  []int
	score int
}

var playerList map[string]player

func Serve() {
	r := mux.NewRouter()
	r.Methods("GET").Path("/").HandlerFunc(redirectBase)

	http.Handle("/ui/", http.StripPrefix("/ui/", http.FileServer(http.Dir("ui"))))
	//http.HandleFunc("/ws", wsHandler)
	http.Handle("/", r)
	bind()

}

func redirectBase(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/ui", http.StatusFound)
}

func bind() {
	port := "8080"

	fmt.Printf("Starting web ui on http://localhost:%s", port)
	if err := ListenAndServe(":" + port); err != nil {
		panic(err)
	}
}

var ListenAndServe = func(bind string) error {
	return http.ListenAndServe(bind, nil)
}
