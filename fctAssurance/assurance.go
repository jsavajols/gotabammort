package assurance

import (
	"encoding/json"
	"fmt"

	c "../const"
)

// Assurance Ajout des mensualités
func Assurance(ta []byte) []byte {
	var retourMoteur c.RetourMoteur
	json.Unmarshal(ta, &retourMoteur)
	fmt.Print(retourMoteur.Capital)
	for m := 0; m < retourMoteur.Duree; m++ {
		fmt.Print(retourMoteur.Periodes[m])
	}

	return ta
}
