package plugnotas

type CorrecaoStatus struct {
	Message *string       `json:"message,omitempty"`
	Data    *CorrecaoData `json:"data,omitempty"`
}

type CorrecaoData struct {
	Date            *string `json:"date,omitempty"`
	Correcao        *string `json:"correcao,omitempty"`
	NumeroSequencia *string `json:"numeroSequencia,omitempty"`
	Status          *string `json:"status,omitempty"`
	RespostaSefaz   *string `json:"respostaSefaz,omitempty"`
	PDF             *string `json:"pdf,omitempty"`
	XML             *string `json:"xml,omitempty"`
}
