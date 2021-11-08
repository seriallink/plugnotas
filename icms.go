package plugnotas

type ICMS struct {
	Origem                 *string                    `json:"origem,omitempty"`
	CST                    *string                    `json:"cst,omitempty"`
	Aliquota               *float64                   `json:"aliquota,omitempty"`
	Valor                  *float64                   `json:"valor,omitempty"`
	Desoneracao            *Desoneracao               `json:"desoneracao,omitempty"`
	BaseCalculo            *BaseCalculoObj            `json:"baseCalculoIcms,omitempty"`
	FundoCombatePobreza    *FundoCombatePobreza       `json:"fundoCombatePobreza,omitempty"`
	Efetivo                *Efetivo                   `json:"efetivo,omitempty"`
	OperacaoPropria        *OperacaoPropria           `json:"operacaoPropria,omitempty"`
	SubstituicaoTributaria *SubstituicaoTributariaObj `json:"substituicaoTributaria,omitempty"`
	Diferimento            *Diferimento               `json:"diferimento,omitempty"`
	CreditoSimplesNacional *CreditoSimplesNacional    `json:"creditoSimplesNacional,omitempty"`
}
