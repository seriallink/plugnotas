package plugnotas

type Total struct {
	BaseCalculoIcms          *float64      `json:"baseCalculoIcms,omitempty"`
	ValorIcms                *float64      `json:"valorIcms,omitempty"`
	ValorFcp                 *float64      `json:"valorFcp,omitempty"`
	ValorFcpSt               *float64      `json:"valorFcpSt,omitempty"`
	BaseCalculoIrrf          *float64      `json:"baseCalculoIrrf,omitempty"`
	ValorIrrfRetido          *float64      `json:"valorIrrfRetido,omitempty"`
	ValorFcpStRetido         *float64      `json:"valorFcpStRetido,omitempty"`
	ValorIcmsDesonerado      *float64      `json:"valorIcmsDesonerado,omitempty"`
	BaseCalculoIcmsSt        *float64      `json:"baseCalculoIcmsSt,omitempty"`
	ValorIcmsSt              *float64      `json:"valorIcmsSt,omitempty"`
	ValorProdutosServicos    *float64      `json:"valorProdutosServicos,omitempty"`
	ValorFrete               *float64      `json:"valorFrete,omitempty"`
	ValorSeguro              *float64      `json:"valorSeguro,omitempty"`
	ValorDesconto            *float64      `json:"valorDesconto,omitempty"`
	ValorIi                  *float64      `json:"valorIi,omitempty"`
	ValorIpi                 *float64      `json:"valorIpi,omitempty"`
	ValorIpiDevolvido        *float64      `json:"valorIpiDevolvido,omitempty"`
	ValorPis                 *float64      `json:"valorPis,omitempty"`
	ValorCofins              *float64      `json:"valorCofins,omitempty"`
	ValorOutros              *float64      `json:"valorOutros,omitempty"`
	ValorNfe                 *float64      `json:"valorNfe,omitempty"`
	ValorAproximadoTributos  *float64      `json:"valorAproximadoTributos,omitempty"`
	ValorPisRetido           *float64      `json:"valorPisRetido,omitempty"`
	ValorCofinsRetido        *float64      `json:"valorCofinsRetido,omitempty"`
	ValorCsllRetido          *float64      `json:"valorCsllRetido,omitempty"`
	BaseCalculoRetencao      *float64      `json:"baseCalculoRetencao,omitempty"`
	ValorPrevidenciaRetido   *float64      `json:"valorPrevidenciaRetido,omitempty"`
	ValorIcmsEstadoDestino   *float64      `json:"valorIcmsEstadoDestino,omitempty"`
	ValorIcmsEstadoRemetente *float64      `json:"valorIcmsEstadoRemetente,omitempty"`
	Servico                  *ServicoTotal `json:"servico,omitempty"`
}
