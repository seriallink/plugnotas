package plugnotas

type Parcela struct {
	TipoPagamento  *float64 `json:"tipoPagamento,omitempty"`
	Numero         *float64 `json:"numero,omitempty"`
	DataVencimento *string  `json:"dataVencimento,omitempty"`
	Valor          *float64 `json:"valor,omitempty"`
}
