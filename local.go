package plugnotas

type Local struct {
	CpfCnpj           string    `json:"cpfCnpj,omitempty"`
	RazaoSocial       string    `json:"razaoSocial,omitempty"`
	InscricaoEstadual string    `json:"inscricaoEstadual,omitempty"`
	Email             string    `json:"email,omitempty"`
	Telefone          *Telefone `json:"telefone,omitempty"`
	Endereco          *Endereco `json:"endereco,omitempty"`
}
