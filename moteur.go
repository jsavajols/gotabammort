package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	c "./const"
	a "./fctAssurance"
	t "./fctmoteur"
)

/*
func main() {
	// go run moteur.go '{"DateEffetGar": "2020-08-01", "Capital": 200000, "Duree": 4, "TxInteret": 1, "PeriodeCalcul": 12, "DtNais": "1962-05-24", "Csp": 1}'
	var params c.Params
	json.Unmarshal([]byte(os.Args[1]), &params)

	ta := t.Ta(params)
	//fmt.Printf(string(ta))
	ass := a.Assurance(ta, params)
	fmt.Printf(string(ass))
}
*/

// CallMoteur = CallMoteur
func CallMoteur(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var params c.Params
		json.NewDecoder(r.Body).Decode(&params)
		fmt.Print(params)
		ta := t.Ta(params)
		//fmt.Printf(string(ta))
		ass := a.Assurance(ta, params)
		retTa := t.AjusteTa(ass, params)
		w.Header().Set("Content-Type", "application/json")
		w.Write(retTa)
	}
}

func main() {
	// curl -i -X POST -H "Content-Type: application/json" -d '{"DateEffetGar": "2020-08-01", "Capital": 200000, "Duree": 4, "TxInteret": 1, "PeriodeCalcul": 12, "DtNais": "1962-05-24", "Csp": 1}' "http://127.0.0.1:8080/"
	http.HandleFunc("/", CallMoteur)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
