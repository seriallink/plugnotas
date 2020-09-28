package plugnotas

type Efetivo struct {
	Valor             *float64          `json:"valor,omitempty"`
	Aliquota          *float64          `json:"aliquota,omitempty"`
	PercentualReducao *float64          `json:"percentualReducao,omitempty"`
	BaseCalculo       *BaseCalculoValor `json:"baseCalculo,omitempty"`
}
