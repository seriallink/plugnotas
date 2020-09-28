package plugnotas

type Medicamento struct {
	CodigoAnvisa         *string `json:"codigoAnvisa,omitempty"`
	MotivoInsencaoAnvisa *string `json:"motivoInsencaoAnvisa,omitempty"`
	ValorMaximo          *string `json:"valorMaximo,omitempty"`
}
