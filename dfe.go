package plugnotas

type DFeConfig struct {
	Ativo *bool `json:"ativo,omitempty"`
}

func (c *Client) ConsultarDFe(cpfCnpj string, manifestada bool) (*NFeResumo, error) {

	var (
		err  error
		nfes []NFeResumo
	)

	params := &Params{
		"cpfCnpj":     cpfCnpj,
		"manifestada": manifestada,
	}

	if err = c.Get("nfe/destinada", params, nil, &nfes); err == nil && len(nfes) == 1 {
		return &nfes[0], err
	}

	return nil, err

}
