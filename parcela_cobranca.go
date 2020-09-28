package plugnotas

type ParcelaCobranca struct {
	Numero         *string  `json:"numero,omitempty"`
	DataVencimento *string  `json:"dataVencimento,omitempty"`
	Valor          *float64 `json:"valor,omitempty"`
}
