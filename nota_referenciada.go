package plugnotas

type NotaReferenciada struct {
	NFe           []NFeChave      `json:"nfe,omitempty"`
	NFePapel      []NFePapel      `json:"nfePapel,omitempty"`
	ProdutorRural []ProdutorRural `json:"produtorRural,omitempty"`
	CupomFiscal   []CupomFiscal   `json:"cupomFiscal,omitempty"`
}
