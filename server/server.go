package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	PortVar = "VCAP_APP_PORT"
)

const width int = 100
const height int = 100

type player struct {
	name  string
	snake []string
}

var playerList []player

func Serve() {
	r := mux.NewRouter()
	r.Methods("GET").Path("/").HandlerFunc(redirectBase)
	r.Methods("GET").Path("/update").HandlerFunc(updateHandler)

	http.Handle("/ui/", http.StripPrefix("/ui/", http.FileServer(http.Dir("ui"))))

	http.Handle("/", r)
	bind()

}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	snake := []string{"1,1", "2,1", "3,1"}
	p := &player{"simon", snake}
	d, _ := json.Marshal(p)
	//fmt.Fprintf(w, string(d))
	fmt.Println(d)
	fmt.Fprintf(w, "", d)
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
