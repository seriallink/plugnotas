package plugnotas

type SubstituicaoTributaria struct {
	BaseCalculo       *float64 `json:"baseCalculo,omitempty"`
	Aliquota          *float64 `json:"aliquota,omitempty"`
	QuantidadeVendida *float64 `json:"quantidadeVendida,omitempty"`
	AliquotaReais     *float64 `json:"aliquotaReais,omitempty"`
	Valor             *float64 `json:"valor,omitempty"`
}
