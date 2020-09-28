package plugnotas

type NFePapel struct {
	Estado      *string `json:"estado,omitempty"`
	DataEmissao *string `json:"dataEmissao,omitempty"`
	CpfCnpj     *string `json:"cpfCnpj,omitempty"`
	Modelo      *string `json:"modelo,omitempty"`
	Serie       *string `json:"serie,omitempty"`
	Numero      *string `json:"numero,omitempty"`
}
