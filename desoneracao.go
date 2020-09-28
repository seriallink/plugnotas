package plugnotas

type Desoneracao struct {
	Valor  *float64 `json:"valor,omitempty"`
	Motivo *string  `json:"motivo,omitempty"`
}
