package plugnotas

type ISS struct {
	TipoTributacao    *int     `json:"tipoTributacao,omitempty"`
	Exigibilidade     *int     `json:"exigibilidade,omitempty"`
	Retido            *bool    `json:"retido,omitempty"`
	Aliquota          *float64 `json:"aliquota,omitempty"`
	Valor             *float64 `json:"valor,omitempty"`
	ValorRetido       *float64 `json:"valorRetido,omitempty"`
	ProcessoSuspensao *string  `json:"processo_suspensao,omitempty"`
}
