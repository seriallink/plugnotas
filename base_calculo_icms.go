package plugnotas

type BaseCalculoIcms struct {
	PercentualReducao      *float64 `json:"percentualReducao,omitempty"`
	ModalidadeDeterminacao *float64 `json:"modalidadeDeterminacao,omitempty"`
	Valor                  *float64 `json:"valor,omitempty"`
}
