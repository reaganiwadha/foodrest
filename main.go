package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

//Food - A struct of food
type Food struct {
	ID    int
	Name  string
	Price float32
}

var deletedCount = 0

var foods = []Food{
	Food{
		ID:    0,
		Name:  "Pasta",
		Price: 5,
	},
	Food{
		ID:    1,
		Name:  "Burger",
		Price: 3,
	},
}

//RemoveIndex - A method to remove an index in an array
func RemoveIndex(s []Food, index int) []Food {
	return append(s[:index], s[index+1:]...)
}

//RespondWithJSON - A method to respond with json data
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

//RespondWithOK - A method to return a 200 code
func RespondWithOK(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
}

//RespondWithError - Throws an error
func RespondWithError(w http.ResponseWriter, code int, msg string) {
	RespondWithJSON(w, code, map[string]string{"error": msg})
}

// CreateFoodEndpoint - An endpoint to add food
func CreateFoodEndpoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var food Food
	if err := json.NewDecoder(r.Body).Decode(&food); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
	}

	food.ID = len(foods) + deletedCount

	foods = append(foods, food)
	RespondWithOK(w)
}

// AllFoodsEndpoint - An endpoint to return all the foods
func AllFoodsEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("AllFoodsEndpoint is accessed")
	RespondWithJSON(w, http.StatusOK, foods)
}

// GetFoodByIDEndpoint - An endpoint to return a food based on its id
func GetFoodByIDEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
	}

	var index int
	for i := range foods {
		if foods[i].ID == id {
			index = i
		}
	}

	var food Food
	if err := json.NewDecoder(r.Body).Decode(&food); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
	}

	RespondWithJSON(w, http.StatusOK, foods[index])
}

// UpdateFoodByIDEndpoint - An endpoint to update a food based on its id
func UpdateFoodByIDEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
	}

	var index int
	for i := range foods {
		if foods[i].ID == id {
			index = i
		}
	}

	var food Food
	if err := json.NewDecoder(r.Body).Decode(&food); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
	}

	food.ID = foods[index].ID

	foods[index] = food
}

// DeleteFoodByIDEndpoint - An endpoint to delete food based on its id
func DeleteFoodByIDEndpoint(w http.ResponseWriter, r *http.Request) {
	deletedCount++
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
	}

	var index int

	for i := range foods {
		if foods[i].ID == id {
			index = i
		}
	}

	foods = RemoveIndex(foods, index)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := mux.NewRouter()
	router.HandleFunc("/foods", CreateFoodEndpoint).Methods("POST")
	router.HandleFunc("/foods", AllFoodsEndpoint).Methods("GET")
	router.HandleFunc("/foods/{id}", GetFoodByIDEndpoint).Methods("GET")
	router.HandleFunc("/foods/{id}", UpdateFoodByIDEndpoint).Methods("PUT")
	router.HandleFunc("/foods/{id}", DeleteFoodByIDEndpoint).Methods("DELETE")

	fmt.Println("Starting Foodrest...")
	http.Handle("/", router)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)

	fmt.Println("Running on port " + os.Getenv("PORT"))
}
