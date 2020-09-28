package plugnotas

type Combustivel struct {
	CodigoAnp                      *string     `json:"codigoAnp,omitempty"`
	DescricaoAnp                   *string     `json:"descricaoAnp,omitempty"`
	CodigoAutorizacao              *string     `json:"codigoAutorizacao,omitempty"`
	EstadoConsumo                  *string     `json:"estadoConsumo,omitempty"`
	PercentualGlp                  *float64    `json:"percentualGlp,omitempty"`
	PercentualGnn                  *float64    `json:"percentualGnn,omitempty"`
	PercentualGni                  *float64    `json:"percentualGni,omitempty"`
	ValorPartida                   *float64    `json:"valorPartida,omitempty"`
	FaturamentoTemperaturaAmbiente *float64    `json:"faturamentoTemperaturaAmbiente,omitempty"`
	CIDE                           *CIDE       `json:"cide,omitempty"`
	Encerrante                     *Encerrante `json:"encerrante,omitempty"`
}
