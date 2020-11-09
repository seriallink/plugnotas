package plugnotas

type NFSe struct {
	IdIntegracao              *string          `json:"idIntegracao,omitempty"`
	IdNotaSubstituida         *string          `json:"idNotaSubstituida,omitempty"`
	Descricao                 *string          `json:"descricao,omitempty"`
	Status                    *string          `json:"status,omitempty"`
	InformacoesComplementares *string          `json:"informacoesComplementares,omitempty"`
	NaturezaTributacao        *int             `json:"naturezaTributacao,omitempty"`
	BaseCalculoIrrf           *float64         `json:"baseCalculoIrrf,omitempty"`
	ValorIrrfRetido           *float64         `json:"valorIrrfRetido,omitempty"`
	EnviarEmail               *bool            `json:"enviarEmail,omitempty"`
	Substituicao              *bool            `json:"substituicao,omitempty"`
	Servico                   []Servico        `json:"servico,omitempty"`
	Parcelas                  []Parcela        `json:"parcelas,omitempty"`
	Prestador                 *Prestador       `json:"prestador,omitempty"`
	Tomador                   *Tomador         `json:"tomador,omitempty"`
	Rps                       *Rps             `json:"rps,omitempty"`
	CidadePrestacao           *CidadePrestacao `json:"cidadePrestacao,omitempty"`
	CargaTributaria           *CargaTributaria `json:"cargaTributaria,omitempty"`
	Impressao                 *Impressao       `json:"impressao,omitempty"`
	CamposExtras              *CamposExtras    `json:"camposExtras,omitempty"`
}

type NFSeObject struct {
	Ativo        *bool       `json:"ativo,omitempty"`
	TipoContrato *int        `json:"tipoContrato,omitempty"`
	Config       *NFSeConfig `json:"config,omitempty"`
}

type NFSeConfig struct {
	Producao   *bool       `json:"producao,omitempty"`
	Rps        *RpsConfig  `json:"rps,omitempty"`
	Prefeitura *Prefeitura `json:"prefeitura,omitempty"`
	Email      *Email      `json:"email,omitempty"`
}

func (c *Client) EmitirNFSe(nfses []NFSe) (*NfResponse, error) {
	resp := new(NfResponse)
	err := c.Post("nfse/", nfses, nil, resp)
	return resp, err
}

func (c *Client) ConsultarNFSe(id string) (*NFSe, error) {
	nfse := new(NFSe)
	err := c.Get("nfse/"+id, nil, nil, nfse)
	return nfse, err
}

func (c *Client) CancelarNFSe(id string) (*NfResponse, error) {
	resp := new(NfResponse)
	err := c.Post("nfse/cancelar/"+id, nil, nil, resp)
	return resp, err
}

func (c *Client) DownloadNFSePDF(id string) (pdf []byte, err error) {
	pdf = make([]byte, 0)
	err = c.Get("nfse/pdf/"+id, nil, nil, &pdf)
	return
}

func (c *Client) DownloadNFSeXML(id string) (xml string, err error) {
	err = c.Get("nfse/xml/"+id, nil, nil, &xml)
	return
}
