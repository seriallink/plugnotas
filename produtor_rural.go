package plugnotas

type ProdutorRural struct {
	Estado            *string `json:"estado,omitempty"`
	DataEmissao       *string `json:"dataEmissao,omitempty"`
	CpfCnpj           *string `json:"cpfCnpj,omitempty"`
	InscricaoEstadual *string `json:"inscricaoEstadual,omitempty"`
	Modelo            *string `json:"modelo,omitempty"`
	Serie             *string `json:"serie,omitempty"`
	Numero            *string `json:"numero,omitempty"`
	ChaveCte          *string `json:"chaveCte,omitempty"`
}
