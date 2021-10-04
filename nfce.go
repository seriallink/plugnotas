package plugnotas

type NFCeObject struct {
	Ativo        *bool       `json:"ativo,omitempty"`
	TipoContrato *int        `json:"tipoContrato,omitempty"`
	Config       *NFCeConfig `json:"config,omitempty"`
}

type NFCeConfig struct {
	Producao    *bool        `json:"producao,omitempty"`
	Email       *Email       `json:"email,omitempty"`
	Sefaz       *Sefaz       `json:"sefaz,omitempty"`
	Numeracao   []Numeracao  `json:"numeracao,omitempty"`
	Integracoes *Integracoes `json:"integracoes,omitempty"`
}
