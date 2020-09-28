package plugnotas

type FundoCombatePobreza struct {
	Aliquota    *float64          `json:"aliquota,omitempty"`
	Valor       *float64          `json:"valor,omitempty"`
	BaseCalculo *BaseCalculoValor `json:"baseCalculo,omitempty"`
}
