package plugnotas

type Imposto struct {
	Aliquota *float64 `json:"aliquota,omitempty"`
	Valor    *float64 `json:"valor,omitempty"`
}
