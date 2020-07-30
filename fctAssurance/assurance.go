package assurance

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"time"

	c "../const"
	d "../fctDates"
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
	dateSituation := d.StringToDate(params.DateEffetGar)

	// Charge les Tarifs
	tarifs := Tarifs
	data, _ := ioutil.ReadFile("./tarifs/tarifsRAC.json")
	_ = json.Unmarshal(data, &tarifs)
	for m := 0; m < retourMoteur.Duree; m++ {
		// fmt.Print(retourMoteur.Periodes[m])
		ageAssure := age(dateSituation, params, "ann")
		tarif := getTarif(ageAssure, params.Csp, tarifs)
		retourMoteur.Periodes[m].Assurance = c.PeriodeAssurance{
			Age: ageAssure,
			Csp: params.Csp,
			Dc:  (retourMoteur.Periodes[m].Remboursement * tarif.Dc) / 100,
			It:  (retourMoteur.Periodes[m].Remboursement * tarif.It) / 100,
			Ipp: (retourMoteur.Periodes[m].Remboursement * tarif.Ipp) / 100,
			Exo: (retourMoteur.Periodes[m].Remboursement * tarif.Exo) / 100,
		}
		dateSituation = dateSituation.AddDate(0, 1, 0)
	}
	jsonData, _ := json.Marshal(retourMoteur)

	return jsonData
}

func age(dateSit time.Time, params c.Params, typeCalcul string) int {
	var ageCalcul int
	anneeGar := dateSit.Year()
	fmt.Print(dateSit)
	// Calcul en milesime
	if typeCalcul == "mil" {
		anneeNais, _ := strconv.ParseInt(params.DtNais[0:4], 10, 64)
		ageCalcul = int(anneeGar - int(anneeNais))
	}
	if typeCalcul == "ann" {
		dateNaisAss := d.StringToDate(params.DtNais)
		days := dateSit.Sub(dateNaisAss).Hours() / 24
		ageCalcul = int(math.Trunc(days / 360))
	}
	// Calcul en jours
	return ageCalcul
}

func getTarif(age int, csp int, tarifs []Tarif) Tarif {

	var retTarif Tarif
	for t := 0; t < len(tarifs); t++ {
		if tarifs[t].Age == age && tarifs[t].Csp == csp {
			retTarif = tarifs[t]
			break
		}
	}
	return retTarif
}
