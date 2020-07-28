package main

import (
	"encoding/json"
	"fmt"
	"math"
)

// Periode Structure
type Periode struct {
	TypePeriode   string
	NumPeriode    int
	CrdDebut      float64
	CrdFin        float64
	Remboursement float64
	Interets      float64
}

// Periodes = Periodes
var Periodes []Periode

// RetourMoteur = RetourMoteur
type RetourMoteur struct {
	Capital    float64
	Duree      int
	txInteret  float64
	Iterations int
	Periodes   []Periode
}

func moteur(capital float64, duree int, txInteret float64, detail string) []byte {
	var interets float64
	var mensu float64
	var crdDebutMen = capital
	var crdDebutAn = capital
	var crdFin float64
	var annee = 1
	var cumulInterets float64
	txInteret = (txInteret / 100) / 12
	mensu = (capital * txInteret) / (1 - math.Pow((1+(txInteret)), -float64(duree)))

	for m := 1; m <= duree; m++ {
		interets = crdDebutMen * txInteret
		crdFin = crdDebutMen - mensu + interets

		if detail == "m" {
			Periodes = append(Periodes, Periode{
				"Mois",
				m,
				crdDebutMen,
				crdFin,
				mensu,
				interets,
			})
		} else {
			var forTest float64
			forTest = float64(m) / 12
			if forTest == math.Trunc(forTest) {
				// fmt.Printf("*****" + string(forTest))
				Periodes = append(Periodes, Periode{
					"Annee",
					annee,
					crdDebutAn,
					crdFin,
					mensu,
					cumulInterets,
				})
				cumulInterets = 0
				crdDebutAn = crdFin
				annee++
			}
		}
		crdDebutMen = crdFin
		cumulInterets = cumulInterets + interets

	}
	var iterations int
	if detail == "m" {
		iterations = duree
	} else {
		iterations = annee - 1
	}
	var retourMoteur = RetourMoteur{
		capital,
		duree,
		txInteret,
		iterations,
		Periodes,
	}
	jsonData, _ := json.Marshal(retourMoteur)
	return jsonData
}

func main() {
	fmt.Printf(string(moteur(200000, 300, 1, "m")))
}
