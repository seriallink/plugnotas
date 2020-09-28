package plugnotas

type Destinatario struct {
	CpfCnpj            *string   `json:"cpfCnpj,omitempty,omitempty"`
	InscricaoMunicipal *string   `json:"inscricaoMunicipal,omitempty"`
	InscricaoEstadual  *string   `json:"inscricaoEstadual,omitempty"`
	RazaoSocial        *string   `json:"razaoSocial,omitempty"`
	NomeFantasia       *string   `json:"nomeFantasia,omitempty"`
	OrgaoPublico       *bool     `json:"orgaoPublico,omitempty"`
	Email              *string   `json:"email,omitempty"`
	Endereco           *Endereco `json:"endereco,omitempty"`
	Telefone           *Telefone `json:"telefone,omitempty"`
}
