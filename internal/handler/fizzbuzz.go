package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Riverfount/fizz-buzz-api/internal/service"
)

func FizzBuzzHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	num, err := strconv.Atoi(r.URL.Query().Get("n"))
	if err != nil {
		http.Error(w, "Attribute error", http.StatusBadRequest)
		return
	}

	response := map[string]string{
		"message": service.FizzBuzz(num),
	}

	json.NewEncoder(w).Encode(response)

}
