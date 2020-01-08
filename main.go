package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type Food struct {
	Name  string
	Price float32
}

var foods = []Food{
	Food{
		Name:  "Pasta",
		Price: 3801.2,
	},
	Food{
		Name:  "Burger",
		Price: 3801.2,
	},
}

//RespondWithJSON - A method to respond with json data
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// AllFoodsEndpoint - An endpoint to return all the foods
func AllFoodsEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("AllFoodsEndpoint is accessed")
	RespondWithJSON(w, http.StatusOK, foods)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := mux.NewRouter()

	router.HandleFunc("/foods", AllFoodsEndpoint).Methods("GET")

	fmt.Println("Starting Foodrest...")

	http.Handle("/", router)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)

	fmt.Println("Running on port " + os.Getenv("PORT"))
}
