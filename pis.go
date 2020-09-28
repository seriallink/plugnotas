package plugnotas

type PIS struct {
	CST                    *string                 `json:"cst,omitempty"`
	Aliquota               *float64                `json:"aliquota,omitempty"`
	Valor                  *float64                `json:"valor,omitempty"`
	QuantidadeVendida      *float64                `json:"quantidadeVendida,omitempty"`
	AliquotaReais          *float64                `json:"aliquotaReais,omitempty"`
	BaseCalculo            *BaseCalculo            `json:"baseCalculo,omitempty"`
	SubstituicaoTributaria *SubstituicaoTributaria `json:"substituicaoTributaria,omitempty"`
}
