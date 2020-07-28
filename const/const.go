package constantes

// TIMERCHECK Nombre de secondes pour le check
const TIMERCHECK = 5

// Periode Structure
type Periode struct {
	TypePeriode   string
	NumPeriode    int
	CrdDebut      float64
	CrdFin        float64
	Remboursement float64
	Interets      float64
}

// RetourMoteur = RetourMoteur
type RetourMoteur struct {
	Capital    float64
	Duree      int
	TxInteret  float64
	Iterations int
	Periodes   []Periode
}
