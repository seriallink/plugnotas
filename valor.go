package plugnotas

type Valor struct {
	Servico                 *float64 `json:"servico,omitempty"`
	BaseCalculo             *float64 `json:"baseCalculo,omitempty"`
	Deducoes                *float64 `json:"deducoes,omitempty"`
	DescontoCondicionado    *float64 `json:"descontoCondicionado,omitempty"`
	DescontoIncondicionado  *float64 `json:"descontoIncondicionado,omitempty"`
	Liquido                 *float64 `json:"liquido,omitempty"`
	Unitario                *float64 `json:"unitario,omitempty"`
	ValorAproximadoTributos *float64 `json:"valorAproximadoTributos,omitempty"`
}
