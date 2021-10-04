package plugnotas

type Empresa struct {
	CpfCnpj                  *string     `json:"cpfCnpj,omitempty"`
	InscricaoMunicipal       *string     `json:"inscricaoMunicipal,omitempty"`
	InscricaoEstadual        *string     `json:"inscricaoEstadual,omitempty"`
	RazaoSocial              *string     `json:"razaoSocial,omitempty"`
	NomeFantasia             *string     `json:"nomeFantasia,omitempty"`
	Certificado              *string     `json:"certificado,omitempty"`
	SimplesNacional          *bool       `json:"simplesNacional,omitempty"`
	IncentivoFiscal          *bool       `json:"incentivoFiscal,omitempty"`
	IncentivadorCultural     *bool       `json:"incentivadorCultural,omitempty"`
	RegimeTributario         *int        `json:"regimeTributario,omitempty"`
	RegimeTributarioEspecial *int        `json:"regimeTributarioEspecial,omitempty"`
	Email                    *string     `json:"email,omitempty"`
	Endereco                 *Endereco   `json:"endereco,omitempty"`
	Telefone                 *Telefone   `json:"telefone,omitempty"`
	NFSe                     *NFSeObject `json:"nfse,omitempty"`
	NFe                      *NFeObject  `json:"nfe,omitempty"`
	NFCe                     *NFCeObject `json:"nfce,omitempty"`
}

func (c *Client) ListarEmpresas() (empresas []Empresa, err error) {
	err = c.Get("empresa/", nil, nil, &empresas)
	return
}

func (c *Client) ConsultarEmpresa(cpfCnpj string) (empresa Empresa, err error) {
	err = c.Get("empresa/"+cpfCnpj, nil, nil, &empresa)
	return
}

func (c *Client) CadastrarEmpresa(empresa *Empresa) (err error) {
	err = c.Post("empresa/", empresa, nil, empresa)
	return
}

func (c *Client) AlterarEmpresa(empresa *Empresa) (err error) {
	err = c.Patch("empresa/"+*empresa.CpfCnpj, empresa, nil, empresa)
	return
}
