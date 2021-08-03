package plugnotas

type CIDE struct {
	Valor       *float64 `json:"valor,omitempty"`
	Aliquota    *float64 `json:"aliquota,omitempty"`
	BaseCalculo *float64 `json:"baseCalculo,omitempty"`
}
