package plugnotas

type ServicoTotal struct {
	DataPrestacao               *string  `json:"dataPrestacao,omitempty"`
	BaseCalculo                 *float64 `json:"baseCalculo,omitempty"`
	Valor                       *float64 `json:"valor,omitempty"`
	ValorIss                    *float64 `json:"valorIss,omitempty"`
	ValorPis                    *float64 `json:"valorPis,omitempty"`
	ValorCofins                 *float64 `json:"valorCofins,omitempty"`
	ValorDeducao                *float64 `json:"valorDeducao,omitempty"`
	ValorOutros                 *float64 `json:"valorOutros,omitempty"`
	ValorDescontoCondicionado   *float64 `json:"valorDescontoCondicionado,omitempty"`
	ValorDescontoIncondicionado *float64 `json:"valorDescontoIncondicionado,omitempty"`
	ValorIssRetido              *float64 `json:"valorIssRetido,omitempty"`
	CodigoRegimeEspecial        *float64 `json:"codigoRegimeEspecial,omitempty"`
}
