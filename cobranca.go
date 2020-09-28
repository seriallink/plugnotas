package plugnotas

type Cobranca struct {
	Fatura  *Fatura           `json:"fatura,omitempty"`
	Parcela []ParcelaCobranca `json:"parcela,omitempty"`
}
