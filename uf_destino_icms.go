package plugnotas

type UfDestinoIcms struct {
	Valor       *float64          `json:"valor,omitempty"`
	BaseCalculo *BaseCalculoValor `json:"baseCalculo,omitempty"`
}
