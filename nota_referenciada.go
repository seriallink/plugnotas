package plugnotas

type NotasReferenciada struct {
	NFe           []NotaReferenciadaNFe `json:"nfe,omitempty"`
	NFePapel      []NFePapel            `json:"nfePapel,omitempty"`
	ProdutorRural []ProdutorRural       `json:"produtorRural,omitempty"`
	CupomFiscal   []CupomFiscal         `json:"cupomFiscal,omitempty"`
}

type NotaReferenciadaNFe struct {
	Chave *string `json:"chave,omitempty"`
}
