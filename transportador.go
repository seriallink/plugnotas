package plugnotas

type Transportador struct {
	Cnpj              *string   `json:"cnpj,omitempty"`
	Cpf               *string   `json:"cpf,omitempty"`
	Nome              *string   `json:"nome,omitempty"`
	InscricaoEstadual *string   `json:"inscricaoEstadual,omitempty"`
	Endereco          *Endereco `json:"endereco,omitempty"`
}
