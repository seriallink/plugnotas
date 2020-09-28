package plugnotas

type Retencao struct {
	PIS             *Imposto `json:"pis,omitempty"`
	COFINS          *Imposto `json:"cofins,omitempty"`
	CSLL            *Imposto `json:"csll,omitempty"`
	Inss            *Imposto `json:"inss,omitempty"`
	Irrf            *Imposto `json:"irrf,omitempty"`
	CPP             *Imposto `json:"cpp,omitempty"`
	OutrasRetencoes *float64 `json:"outrasRetencoes,omitempty"`
}
