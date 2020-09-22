package plugnotas

type Nfse struct {
	Ativo        bool        `json:"ativo,omitempty"`
	TipoContrato int         `json:"tipoContrato,omitempty"`
	Config       *NfseConfig `json:"config,omitempty"`
}

type NfseConfig struct {
	Producao   bool        `json:"producao,omitempty"`
	Rps        *Rps        `json:"rps,omitempty"`
	Prefeitura *Prefeitura `json:"prefeitura,omitempty"`
	Email      *Email      `json:"email,omitempty"`
}
