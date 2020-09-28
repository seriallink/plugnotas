package plugnotas

type IPI struct {
	CnpjProdutor             *string       `json:"cnpjProdutor,omitempty"`
	CodigoEnquadramentoLegal *string       `json:"codigoEnquadramentoLegal,omitempty"`
	CST                      *string       `json:"cst,omitempty"`
	BaseCalculo              *float64      `json:"baseCalculo,omitempty"`
	Aliquota                 *float64      `json:"aliquota,omitempty"`
	Valor                    *float64      `json:"valor,omitempty"`
	SeloControle             *SeloControle `json:"seloControle,omitempty"`
	Unidade                  *UnidadeIpi   `json:"unidade,omitempty"`
}
