package plugnotas

type Cartao struct {
	TipoIntegracao    *int    `json:"tipoIntegracao,omitempty"`
	CnpjCredenciadora *string `json:"cnpjCredenciadora,omitempty"`
	BandeiraOperadora *string `json:"bandeiraOperadora,omitempty"`
	NumeroAutorizacao *string `json:"numeroAutorizacao,omitempty"`
}
