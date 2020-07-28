package ta

import (
	"encoding/json"
	"math"

	c "../const"
)

// Periodes = Periodes
var Periodes []c.Periode

// Ta Calcul du Ta
func Ta(capital float64, duree int, txInteret float64, detail string) []byte {
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
			Periodes = append(Periodes, c.Periode{
				"Mois",
				m,
				crdDebutMen,
				crdFin,
				mensu,
				interets,
			})
		} else {
			var periodicite float64
			periodicite = float64(m) / 12
			if periodicite == math.Trunc(periodicite) {
				// fmt.Printf("*****" + string(periodicite))
				Periodes = append(Periodes, c.Periode{
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
	var retourMoteur = c.RetourMoteur{
		capital,
		duree,
		txInteret,
		iterations,
		Periodes,
	}
	jsonData, _ := json.Marshal(retourMoteur)
	return jsonData
}
