package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var captain = []Person{
	{Name: "Zhalgas", Age: 19},
}

var players = []Person{
	{Name: "Zhanelya", Age: 16},
	{Name: "Arman", Age: 18},
	{Name: "Adilkhan", Age: 20},
}

var templates = template.Must(template.ParseFiles("templates/base.html", "templates/home.html"))

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	err := templates.ExecuteTemplate(w, "base.html", nil)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func GetCaptain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(captain)
}

func GetPlayersList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(players)
}

func GetPlayersDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	name := params["name"]

	for _, player := range players {
		if player.Name == name {
			json.NewEncoder(w).Encode(player)
			return
		}
	}

	http.Error(w, fmt.Sprintf("Friend %s not found", name), http.StatusNotFound)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, "App is healthy!\nAuthor: Zhalgas")
}
