package plugnotas

type Endereco struct {
	TipoLogradouro  string `json:"tipoLogradouro,omitempty"`
	Logradouro      string `json:"logradouro,omitempty"`
	Numero          string `json:"numero,omitempty"`
	Complemento     string `json:"complemento,omitempty"`
	TipoBairro      string `json:"tipoBairro,omitempty"`
	Bairro          string `json:"bairro,omitempty"`
	CodigoPais      string `json:"codigoPais,omitempty"`
	DescricaoPais   string `json:"descricaoPais,omitempty"`
	CodigoCidade    string `json:"codigoCidade,omitempty"`
	DescricaoCidade string `json:"descricaoCidade,omitempty"`
	Estado          string `json:"estado,omitempty"`
	Cep             string `json:"cep,omitempty"`
}

type IBGE struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Municipio   string `json:"municipio"`
	UF          string `json:"uf"`
	Codigo      string `json:"ibge"`
}

func (c *Client) ConsultaEndereco(cep string) (ibge IBGE, err error) {
	err = c.Get("cep/"+cep, nil, nil, &ibge)
	return
}
