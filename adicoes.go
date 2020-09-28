package plugnotas

type Adicoes struct {
	Numero           *string `json:"numero,omitempty"`
	NumeroSequencia  *string `json:"numeroSequencia,omitempty"`
	CodigoFabricante *string `json:"codigoFabricante,omitempty"`
	ValorDesconto    *string `json:"valorDesconto,omitempty"`
	NumeroDrawback   *string `json:"numeroDrawback,omitempty"`
}
