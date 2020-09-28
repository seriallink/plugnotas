package plugnotas

type CupomFiscal struct {
	Modelo               *string `json:"modelo,omitempty"`
	NumeroOrdemsequencia *string `json:"numeroOrdemsequencia,omitempty"`
	NumeroContador       *string `json:"numeroContador,omitempty"`
}
