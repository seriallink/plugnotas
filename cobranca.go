package plugnotas

type Cobranca struct {
	Fatura        *Fatura           `json:"fatura,omitempty"`
	Numero        *string           `json:"numero,omitempty"`
	ValorTotal    *float64          `json:"valorTotal,omitempty"`
	ValorDesconto *float64          `json:"valorDesconto,omitempty"`
	ValorLiquido  *float64          `json:"valorLiquido,omitempty"`
	Parcela       []ParcelaCobranca `json:"parcela,omitempty"`
}
