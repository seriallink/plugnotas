package plugnotas

type Empresa struct {
	CpfCnpj                  string    `json:"cpfCnpj,omitempty"`
	InscricaoMunicipal       string    `json:"inscricaoMunicipal,omitempty"`
	InscricaoEstadual        string    `json:"inscricaoEstadual,omitempty"`
	RazaoSocial              string    `json:"razaoSocial,omitempty"`
	NomeFantasia             string    `json:"nomeFantasia,omitempty"`
	Certificado              string    `json:"certificado,omitempty"`
	SimplesNacional          bool      `json:"simplesNacional,omitempty"`
	RegimeTributario         int       `json:"regimeTributario,omitempty"`
	IncentivoFiscal          bool      `json:"incentivoFiscal,omitempty"`
	IncentivadorCultural     bool      `json:"incentivadorCultural,omitempty"`
	RegimeTributarioEspecial int       `json:"regimeTributarioEspecial,omitempty"`
	Email                    string    `json:"email,omitempty"`
	Endereco                 *Endereco `json:"endereco,omitempty"`
	Telefone                 *Telefone `json:"telefone,omitempty"`
	Nfse                     *Nfse     `json:"nfse,omitempty"`
	Nfe                      *Nfe      `json:"nfe,omitempty"`
}

func (c *Client) ListaEmpresas() (empresas []Empresa, err error) {
	err = c.Get("empresa/", nil, nil, &empresas)
	return
}

func (c *Client) ConsultaEmpresa(cpfCnpj string) (empresa Empresa, err error) {
	err = c.Get("empresa/"+cpfCnpj, nil, nil, &empresa)
	return
}

func (c *Client) CadastrarEmpresa(empresa *Empresa) (err error) {
	err = c.Post("empresa/", empresa, nil, &empresa)
	return
}
