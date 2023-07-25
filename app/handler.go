package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Bank struct {
	Name   string `json:"name"`
	Number string `json:"number"`
	Client string `json:"client"`
}

func CreateBank(w http.ResponseWriter, r *http.Request) {

}

func GetBank(w http.ResponseWriter, r *http.Request) {
	bank := []Bank{
		{Name: "Itau", Number: "0232", Client: "Jennifer"},
		{Name: "Nubank", Number: "9087", Client: "Aurora"},
		{Name: "PicPay", Number: "6709", Client: "Willow"},
	}

	w.Header().Add("Content-Type", "application/json")

	json.NewEncoder(w).Encode(bank)
}

func GetBankById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["id"])
}
