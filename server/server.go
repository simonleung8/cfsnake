package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const (
	PortVar = "VCAP_APP_PORT"
)

const width int = 100
const height int = 100

type Player struct {
	Name  string
	Snake []string
}

type playerList []Player

var gameData playerList

func init() {
	snake1 := []string{"1,1", "2,1", "3,1"}
	snake2 := []string{"24,1", "25,1", "26,1", "26,2"}
	gameData = playerList{Player{"simon", snake1}, Player{"dan", snake2}}
}

func Serve() {
	r := mux.NewRouter()
	r.Methods("GET").Path("/").HandlerFunc(redirectBase)

	r.Methods("GET").Path("/update").HandlerFunc(updateHandler)
	r.Methods("GET").Path("/left").HandlerFunc(left)
	r.Methods("GET").Path("/up").HandlerFunc(up)
	r.Methods("GET").Path("/right").HandlerFunc(right)
	r.Methods("GET").Path("/down").HandlerFunc(down)

	http.Handle("/ui/", http.StripPrefix("/ui/", http.FileServer(http.Dir("ui"))))

	http.Handle("/", r)

	bind()
}

func goForward() {
	ticker := time.NewTicker(1 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				calcualteSteps()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

func calcualteSteps() {
	fmt.Println(len(gameData))
	for _, player := range gameData {
		for _, snake := range player.Snake {
			//go forward
			fmt.Println(snake)
		}
	}
}

func updateHandler(w http.ResponseWriter, r *http.Request) {

	d, _ := json.Marshal(gameData)
	fmt.Fprintf(w, string(d))
}

func left(w http.ResponseWriter, r *http.Request) {
	fmt.Println("left")
}
func up(w http.ResponseWriter, r *http.Request) {
	fmt.Println("up")
}
func right(w http.ResponseWriter, r *http.Request) {
	fmt.Println("right")
}
func down(w http.ResponseWriter, r *http.Request) {
	fmt.Println("down")
}

func redirectBase(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/ui", http.StatusFound)
}

func bind() {
	port := "8080"

	go goForward()
	fmt.Printf("Starting server on http://localhost:%s", port)
	if err := ListenAndServe(":" + port); err != nil {
		panic(err)
	}
}

var ListenAndServe = func(bind string) error {
	return http.ListenAndServe(bind, nil)
}
