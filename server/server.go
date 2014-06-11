package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	. "github.com/simonleung8/cfsnake/game"
)

const (
	PortVar = "VCAP_APP_PORT"
)

var snakeGame Game

func init() {
	snakeGame.New()
}

func hello(res http.ResponseWriter, r *http.Request) {
	env := os.Environ()
	fmt.Fprintln(res, "ENV:\n")
	for _, e := range env {
		fmt.Fprintln(res, e)
	}
}

func Serve() {
	r := mux.NewRouter()
	r.Methods("GET").Path("/").HandlerFunc(redirectBase)

	r.Methods("GET").Path("/info").HandlerFunc(hello)
	r.Methods("GET").Path("/newPlayer").HandlerFunc(newPlayer)
	r.Methods("GET").Path("/update").HandlerFunc(updateHandler)
	r.Methods("GET").Path("/left/{token}").HandlerFunc(left)
	r.Methods("GET").Path("/up/{token}").HandlerFunc(up)
	r.Methods("GET").Path("/right/{token}").HandlerFunc(right)
	r.Methods("GET").Path("/down/{token}").HandlerFunc(down)

	http.Handle("/ui/", http.StripPrefix("/ui/", http.FileServer(http.Dir("ui"))))
	http.Handle("/", r)

	bind()
}

func newPlayer(w http.ResponseWriter, r *http.Request) {
	token := snakeGame.NewPlayer()
	snakeGame.Start()
	fmt.Fprintf(w, token)
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	d, _ := json.Marshal(snakeGame.GetPlayersMap())
	fmt.Fprintf(w, string(d))
}

func left(w http.ResponseWriter, r *http.Request) {
	token := mux.Vars(r)["token"]
	if snakeGame.Direction(token) != 1 {
		snakeGame.SetDirection(token, 3)
	}
}
func up(w http.ResponseWriter, r *http.Request) {
	token := mux.Vars(r)["token"]
	if snakeGame.Direction(token) != 2 {
		snakeGame.SetDirection(token, 4)
	}
}
func right(w http.ResponseWriter, r *http.Request) {
	token := mux.Vars(r)["token"]
	if snakeGame.Direction(token) != 3 {
		snakeGame.SetDirection(token, 1)
	}
}
func down(w http.ResponseWriter, r *http.Request) {
	token := mux.Vars(r)["token"]
	if snakeGame.Direction(token) != 4 {
		snakeGame.SetDirection(token, 2)
	}
}

func redirectBase(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/ui", http.StatusFound)
}

func bind() {
	port := "8080"

	fmt.Printf("Starting server on http://localhost:%s", port)
	if err := ListenAndServe(":" + port); err != nil {
		panic(err)
	}
}

var ListenAndServe = func(bind string) error {
	return http.ListenAndServe(bind, nil)
}
