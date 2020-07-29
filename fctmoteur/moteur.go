package ta

import (
	"encoding/json"
	"math"

	c "../const"
)

// Periodes = Periodes
var Periodes []c.Periode

// Ta Calcul du Ta
func Ta(params c.Params) []byte {
	var interets float64
	var mensu float64
	var crdDebutMen = params.Capital
	var crdDebutAn = params.Capital
	var crdFin float64
	var annee = 1
	var cumulInterets float64
	var messagesErreur []string
	params.TxInteret = (params.TxInteret / 100) / 12
	if params.TxInteret != 0 {
		mensu = (params.Capital * params.TxInteret) / (1 - math.Pow((1+(params.TxInteret)), -float64(params.Duree)))
	} else {
		mensu = params.Capital / float64(params.Duree)
	}
	var periodeAssurance c.PeriodeAssurance

	for m := 1; m <= params.Duree; m++ {
		if params.TxInteret != 0 {
			interets = crdDebutMen * params.TxInteret
		} else {
			interets = 0
		}
		crdFin = crdDebutMen - mensu + interets

		if params.Detail == "m" {
			Periodes = append(Periodes, c.Periode{
				"Mois",
				m,
				crdDebutMen,
				crdFin,
				mensu,
				interets,
				periodeAssurance,
			})
		} else {
			var periodicite float64
			periodicite = float64(m) / 12
			if periodicite == math.Trunc(periodicite) {
				Periodes = append(Periodes, c.Periode{
					"Annee",
					annee,
					crdDebutAn,
					crdFin,
					mensu,
					cumulInterets,
					periodeAssurance,
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
	if params.Detail == "m" {
		iterations = params.Duree
	} else {
		iterations = annee - 1
	}
	var retourMoteur = c.RetourMoteur{
		params.Capital,
		params.Duree,
		params.TxInteret,
		iterations,
		Periodes,
		0,
		0,
		"",
		messagesErreur,
	}
	jsonData, _ := json.Marshal(retourMoteur)
	return jsonData
}
