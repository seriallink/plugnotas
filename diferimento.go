package plugnotas

type Diferimento struct {
	Valor           *float64 `json:"valor,omitempty"`
	Percentual      *float64 `json:"percentual,omitempty"`
	ValorIcmsDevido *float64 `json:"valorIcmsDevido,omitempty"`
}
