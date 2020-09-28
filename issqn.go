package plugnotas

type ISSQN struct {
	Valor                      *float64 `json:"valor,omitempty"`
	Aliquota                   *float64 `json:"aliquota,omitempty"`
	BaseCalculo                *float64 `json:"baseCalculo,omitempty"`
	ValorDeducao               *float64 `json:"valorDeducao,omitempty"`
	ValorOutros                *float64 `json:"valorOutros,omitempty"`
	DescontoIncondicionado     *float64 `json:"descontoIncondicionado,omitempty"`
	DescontoCondicionado       *float64 `json:"descontoCondicionado,omitempty"`
	ValorRetencaoIss           *float64 `json:"valorRetencaoIss,omitempty"`
	CodigoServico              *string  `json:"codigoServico,omitempty"`
	CodigoMunicipalServico     *string  `json:"codigoMunicipalServico,omitempty"`
	CodigoMunicipioFatoGerador *string  `json:"codigoMunicipioFatoGerador,omitempty"`
	CodigoMunicipioIncidencia  *string  `json:"codigoMunicipioIncidencia,omitempty"`
	CodigoExigibilidade        *string  `json:"codigoExigibilidade,omitempty"`
	NumeroProcessoSuspensao    *string  `json:"numeroProcessoSuspensao,omitempty"`
	IncentivoFiscal            *bool    `json:"incentivoFiscal,omitempty"`
}
