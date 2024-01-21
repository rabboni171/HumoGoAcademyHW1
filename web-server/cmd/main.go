package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func calculate(first, second float64, znak string) (result float64) {
	switch znak {
	case "+":
		result = first + second
	case "-":
		result = first - second
	case "*":
		result = first * second
	case "div":
		result = first / second
	}
	return result
}
func Plus(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	first, err := strconv.Atoi(vars["first"])
	if err != nil {
		http.Error(rw, "Invalid value for 'first'", http.StatusBadRequest)
		return
	}

	second, err := strconv.Atoi(vars["second"])
	if err != nil {
		http.Error(rw, "Invalid value for 'second'", http.StatusBadRequest)
		return
	}

	znak := vars["znak"]

	result := calculate(float64(first), float64(second), znak)
	response := fmt.Sprintf("%f", result)
	fmt.Fprint(rw, response)
}

func main() {
	newMux := mux.NewRouter()
	newMux.HandleFunc("/{first}/{znak}/{second}", Plus)

	fmt.Println("Listening on port 8000")
	err := http.ListenAndServe("127.0.0.1:8000", newMux)
	if err != nil {
		fmt.Println(err)
	}
}

// //http://127.0.0.1:8000/2/+/3
