package plugnotas

type UfDestino struct {
	BaseCalculoIcms                       *float64 `json:"baseCalculoIcms,omitempty"`
	BaseCalculoFcp                        *float64 `json:"baseCalculoFcp,omitempty"`
	PercentualIcmsFcp                     *float64 `json:"percentualIcmsFcp,omitempty"`
	AliquotaInterna                       *float64 `json:"aliquotaInterna,omitempty"`
	IcmsRelativoFcp                       *float64 `json:"icmsRelativoFcp,omitempty"`
	IcmsInterestadual                     *float64 `json:"icmsInterestadual,omitempty"`
	AliquotaInterestadual                 *float64 `json:"aliquotaInterestadual,omitempty"`
	PercentualProvisorioIcmsInterestadual *float64 `json:"percentualProvisorioIcmsInterestadual,omitempty"`
}
