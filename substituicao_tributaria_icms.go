package plugnotas

type SubstituicaoTributariaIcms struct {
	UfDestino             *UfDestinoIcms        `json:"ufDestino,omitempty"`
	BaseCalculo           *BaseCalculoIcms      `json:"baseCalculo,omitempty"`
	MargemValorAdicionado *BaseCalculoPercetual `json:"margemValorAdicionado,omitempty"`
	FundoCombatePobreza   *FundoCombatePobreza  `json:"fundoCombatePobreza,omitempty"`
}
