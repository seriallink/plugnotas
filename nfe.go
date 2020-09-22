package plugnotas

type Nfe struct {
	Ativo        bool       `json:"ativo,omitempty"`
	TipoContrato int        `json:"tipoContrato,omitempty"`
	Config       *NfeConfig `json:"config,omitempty"`
}

type NfeConfig struct {
	Producao          bool `json:"producao,omitempty"`
	ImpressaoFcp      bool `json:"impressaoFcp,omitempty"`
	ImpressaoPartilha bool `json:"impressaoPartilha,omitempty"`
	Serie             int  `json:"serie,omitempty"`
	Numero            int  `json:"numero,omitempty"`
}
