package plugnotas

type SubstituicaoTributariaObj struct {
	Aliquota              *float64              `json:"aliquota,omitempty"`
	UFDestino             *UFDestinoObj         `json:"ufDestino,omitempty"`
	BaseCalculo           *BaseCalculoObj       `json:"baseCalculo,omitempty"`
	MargemValorAdicionado *BaseCalculoPercetual `json:"margemValorAdicionado,omitempty"`
	FundoCombatePobreza   *FundoCombatePobreza  `json:"fundoCombatePobreza,omitempty"`
}
