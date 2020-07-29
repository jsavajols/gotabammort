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
	// go run moteur.go '{"Capital": 200000, "Duree": 4, "TxInteret": 1, "Detail": "m", "DtNais": "1962-05-24", "Csp": 1}'
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
		w.Write(ass)
	}
}

func main() {
	http.HandleFunc("/", CallMoteur)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
