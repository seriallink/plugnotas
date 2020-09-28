package plugnotas

type Encerrante struct {
	NumeroBico   *string  `json:"numeroBico,omitempty"`
	NumeroBomba  *string  `json:"numeroBomba,omitempty"`
	NumeroTanque *string  `json:"numeroTanque,omitempty"`
	ValorInicio  *float64 `json:"valorInicio,omitempty"`
	ValorFinal   *float64 `json:"valorFinal,omitempty"`
}
