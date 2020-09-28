package plugnotas

type Prestador struct {
	CpfCnpj                  *string   `json:"cpfCnpj,omitempty,omitempty"`
	InscricaoMunicipal       *string   `json:"inscricaoMunicipal,omitempty"`
	InscricaoEstadual        *string   `json:"inscricaoEstadual,omitempty"`
	RazaoSocial              *string   `json:"razaoSocial,omitempty"`
	NomeFantasia             *string   `json:"nomeFantasia,omitempty"`
	RegimeTributario         *int      `json:"regimeTributario,omitempty"`
	RegimeTributarioEspecial *int      `json:"regimeTributarioEspecial,omitempty"`
	IncentivoFiscal          *bool     `json:"incentivoFiscal,omitempty"`
	SimplesNacional          *bool     `json:"simplesNacional,omitempty"`
	IncentivadorCultural     *bool     `json:"incentivadorCultural,omitempty"`
	Email                    *string   `json:"email,omitempty"`
	Endereco                 *Endereco `json:"endereco,omitempty"`
	Telefone                 *Telefone `json:"telefone,omitempty"`
}
