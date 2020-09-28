package plugnotas

type Emitente struct {
	CpfCnpj                  *string   `json:"cpfCnpj,omitempty"`
	Required                 *string   `json:"required,omitempty"`
	InscricaoMunicipal       *string   `json:"inscricaoMunicipal,omitempty"`
	InscricaoEstadual        *string   `json:"inscricaoEstadual,omitempty"`
	RazaoSocial              *string   `json:"razaoSocial,omitempty"`
	NomeFantasia             *string   `json:"nomeFantasia,omitempty"`
	RegimeTributario         *int      `json:"regimeTributario,omitempty"`
	RegimeTributarioEspecial *int      `json:"regimeTributarioEspecial,omitempty"`
	SimplesNacional          *bool     `json:"simplesNacional,omitempty"`
	IncentivoFiscal          *bool     `json:"incentivoFiscal,omitempty"`
	IncentivadorCultural     *bool     `json:"incentivadorCultural,omitempty"`
	Email                    *string   `json:"email,omitempty"`
	Endereco                 *Endereco `json:"endereco,omitempty"`
	Telefone                 *Telefone `json:"telefone,omitempty"`
}
