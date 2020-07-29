package assurance

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
	"time"

	c "../const"
)

// Tarif = Tarif
type Tarif struct {
	Age int
	Csp int
	Dc  float64
	It  float64
	Ipp float64
	Exo float64
}

// Tarifs = Tarifs
var Tarifs []Tarif

// Assurance Ajout des mensualités
func Assurance(ta []byte, params c.Params) []byte {
	var retourMoteur c.RetourMoteur
	json.Unmarshal(ta, &retourMoteur)
	//fmt.Print(retourMoteur.Capital)
	age := age(params.DtNais, "mil")

	// Charge les Tarifs
	tarifs := Tarifs
	data, _ := ioutil.ReadFile("./tarifs/tarifsRAC.json")
	_ = json.Unmarshal(data, &tarifs)
	for m := 0; m < retourMoteur.Duree; m++ {
		// fmt.Print(retourMoteur.Periodes[m])
		tarif := getTarif(age, params.Csp, tarifs)
		retourMoteur.Periodes[m].Assurance = c.PeriodeAssurance{
			Age: age,
			Csp: params.Csp,
			Dc:  (retourMoteur.Periodes[m].Remboursement * tarif.Dc) / 100,
			It:  (retourMoteur.Periodes[m].Remboursement * tarif.It) / 100,
			Ipp: (retourMoteur.Periodes[m].Remboursement * tarif.Ipp) / 100,
			Exo: (retourMoteur.Periodes[m].Remboursement * tarif.Exo) / 100,
		}
	}
	jsonData, _ := json.Marshal(retourMoteur)

	return jsonData
}

func age(dtNais string, typeCalcul string) int {
	var ageCalcul int
	// Calcul en milesime
	if typeCalcul == "mil" {
		year, _, _ := time.Now().Date()
		anneeNais, _ := strconv.ParseInt(dtNais[0:4], 10, 64)
		ageCalcul = int(int64(year) - anneeNais)
	}
	// Calcul en jours
	return ageCalcul
}

func getTarif(age int, csp int, tarifs []Tarif) Tarif {

	var retTarif Tarif
	for t := 0; t < len(tarifs); t++ {
		if tarifs[t].Age == age && tarifs[t].Csp == csp {
			retTarif = tarifs[t]
		}
	}
	return retTarif
}
