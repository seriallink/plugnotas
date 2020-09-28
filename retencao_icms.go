package plugnotas

type RetencaoIcms struct {
	ValorServico              *float64 `json:"valorServico,omitempty"`
	BaseICMS                  *float64 `json:"baseICMS,omitempty"`
	Aliquota                  *float64 `json:"aliquota,omitempty"`
	ValorICMSRet              *float64 `json:"valorICMSRet,omitempty"`
	CFOP                      *string  `json:"cfop,omitempty"`
	CodigoMunicipioOcorrencia *string  `json:"codigoMunicipioOcorrencia,omitempty"`
}
