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
	var crdFin float64
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

		Periodes = append(Periodes, c.Periode{
			params.PeriodeCalcul,
			m,
			crdDebutMen,
			crdFin,
			mensu,
			interets,
			periodeAssurance,
		})
		crdDebutMen = crdFin
		cumulInterets = cumulInterets + interets

	}
	var iterations int
	iterations = params.Duree
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

// AjusteTa ajuste le TA en fonction de la périodicité
func AjusteTa(ta []byte, params c.Params) []byte {
	var retourMoteur c.RetourMoteur
	json.Unmarshal(ta, &retourMoteur)
	var newPeriodes []c.Periode
	var newPeriodeAss c.PeriodeAssurance
	var cumulInterets float64
	var cumulRemboursement float64
	var crdDebut = params.Capital
	var crdFin float64
	var rupture = 1
	var ruptures = 1
	for m := 0; m < retourMoteur.Duree; m++ {
		// fmt.Print(retourMoteur.Periodes[m])
		cumulInterets = cumulInterets + retourMoteur.Periodes[m].Interets
		cumulRemboursement = cumulRemboursement + retourMoteur.Periodes[m].Remboursement
		newPeriodeAss.Dc = newPeriodeAss.Dc + retourMoteur.Periodes[m].Assurance.Dc
		if rupture == params.PeriodeCalcul {
			crdFin = crdDebut - (cumulRemboursement - cumulInterets)
			newPeriodeAss = retourMoteur.Periodes[m].Assurance
			newPeriodes = append(newPeriodes, c.Periode{
				params.PeriodeCalcul,
				ruptures,
				crdDebut,
				crdFin,
				cumulRemboursement,
				cumulInterets,
				newPeriodeAss,
			})
			ruptures++
			rupture = 0
			cumulRemboursement = 0
			cumulInterets = 0
			crdDebut = retourMoteur.Periodes[m].CrdFin
		}
		rupture++
	}
	var newRetourMoteur = c.RetourMoteur{
		params.Capital,
		params.Duree,
		params.TxInteret,
		params.PeriodeCalcul,
		newPeriodes,
		0,
		0,
		"",
		retourMoteur.MessagesErreur,
	}
	jsonData, _ := json.Marshal(newRetourMoteur)

	return jsonData
}
