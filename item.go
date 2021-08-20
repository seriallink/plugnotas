package plugnotas

type Item struct {
	Codigo                    *string                    `json:"codigo,omitempty"`
	CodigoBarras              *string                    `json:"codigoBarras,omitempty"`
	CodigoBarrasTributavel    *string                    `json:"codigoBarrasTributavel,omitempty"`
	Descricao                 *string                    `json:"descricao,omitempty"`
	NCM                       *string                    `json:"ncm,omitempty"`
	CEST                      *string                    `json:"cest,omitempty"`
	CFOP                      *string                    `json:"cfop,omitempty"`
	NVE                       *string                    `json:"nve,omitempty"`
	IndicadorEscalaRelevante  *string                    `json:"indicadorEscalaRelevante,omitempty"`
	CnpjFabricante            *string                    `json:"cnpjFabricante,omitempty"`
	CodigoBeneficioFiscal     *string                    `json:"codigoBeneficioFiscal,omitempty"`
	ExTipi                    *string                    `json:"exTipi,omitempty"`
	NumeroCompra              *string                    `json:"numeroCompra,omitempty"`
	PedidoCompra              *string                    `json:"pedidoCompra,omitempty"`
	NumeroFci                 *string                    `json:"numeroFci,omitempty"`
	InformacoesComplementares *string                    `json:"informacoesComplementares,omitempty"`
	Valor                     *float64                   `json:"valor,omitempty"`
	ValorFrete                *float64                   `json:"valorFrete,omitempty"`
	ValorSeguro               *float64                   `json:"valorSeguro,omitempty"`
	ValorDesconto             *float64                   `json:"valorDesconto,omitempty"`
	ValorOutros               *float64                   `json:"valorOutros,omitempty"`
	CompoeTotal               *bool                      `json:"compoeTotal,omitempty"`
	Tributos                  *Tributos                  `json:"tributos,omitempty"`
	Importacao                *Importacao                `json:"importacao,omitempty"`
	Compra                    *Compra                    `json:"compra,omitempty"`
	Unidade                   *ComercialTributavelString `json:"unidade,omitempty"`
	Quantidade                *ComercialTributavelNumber `json:"quantidade,omitempty"`
	ValorUnitario             *ComercialTributavelNumber `json:"valorUnitario,omitempty"`
	Combustivel               *Combustivel               `json:"combustivel,omitempty"`
	Medicamentos              []Medicamento              `json:"medicamentos,omitempty"`
	Armamentos                []Armamento                `json:"armamentos,omitempty"`
}

type ItemDFe struct {
	Codigo                    *string                    `json:"codigo,omitempty"`
	CodigoBarras              *string                    `json:"codigoBarras,omitempty"`
	CodigoBarrasTributavel    *string                    `json:"codigoBarrasTributavel,omitempty"`
	Descricao                 *string                    `json:"descricao,omitempty"`
	NCM                       *string                    `json:"ncm,omitempty"`
	CEST                      *string                    `json:"cest,omitempty"`
	CFOP                      *string                    `json:"cfop,omitempty"`
	NVE                       *string                    `json:"nve,omitempty"`
	IndicadorEscalaRelevante  *string                    `json:"indicadorEscalaRelevante,omitempty"`
	CnpjFabricante            *string                    `json:"cnpjFabricante,omitempty"`
	CodigoBeneficioFiscal     *string                    `json:"codigoBeneficioFiscal,omitempty"`
	ExTipi                    *string                    `json:"exTipi,omitempty"`
	NumeroCompra              *string                    `json:"numeroCompra,omitempty"`
	PedidoCompra              *string                    `json:"pedidoCompra,omitempty"`
	NumeroFci                 *string                    `json:"numeroFci,omitempty"`
	InformacoesComplementares *string                    `json:"informacoesComplementares,omitempty"`
	Valor                     *float64                   `json:"valor,omitempty"`
	ValorFrete                *float64                   `json:"valorFrete,omitempty"`
	ValorSeguro               *float64                   `json:"valorSeguro,omitempty"`
	ValorDesconto             *float64                   `json:"valorDesconto,omitempty"`
	ValorOutros               *float64                   `json:"valorOutros,omitempty"`
	CompoeTotal               *bool                      `json:"compoeTotal,omitempty"`
	Tributos                  *Tributos                  `json:"tributos,omitempty"`
	Importacao                []Importacao               `json:"importacao,omitempty"`
	Compra                    *Compra                    `json:"compra,omitempty"`
	Unidade                   *ComercialTributavelString `json:"unidade,omitempty"`
	Quantidade                *ComercialTributavelNumber `json:"quantidade,omitempty"`
	ValorUnitario             *ComercialTributavelNumber `json:"valorUnitario,omitempty"`
	Combustivel               *Combustivel               `json:"combustivel,omitempty"`
	Medicamentos              []Medicamento              `json:"medicamentos,omitempty"`
	Armamentos                []Armamento                `json:"armamentos,omitempty"`
}
