package constantes

// Params = Params
type Params struct {
	Capital   float64
	Duree     int
	TxInteret float64
	Detail    string
	DtNais    string
	Csp       int
}

// PeriodeAssurance = PeriodeAssurance
type PeriodeAssurance struct {
	Age int
	Csp int
	Dc  float64
	It  float64
	Ipp float64
	Exo float64
}

// Periode Structure
type Periode struct {
	TypePeriode   string
	NumPeriode    int
	CrdDebut      float64
	CrdFin        float64
	Remboursement float64
	Interets      float64
	Assurance     PeriodeAssurance
}

// RetourMoteur = RetourMoteur
type RetourMoteur struct {
	Capital        float64
	Duree          int
	TxInteret      float64
	Iterations     int
	Periodes       []Periode
	CodeFin        int
	CodeErreur     int
	MessageFin     string
	MessagesErreur []string
}
