package plugnotas

type Rps struct {
	Competencia    *string `json:"competencia,omitempty"`
	DataEmissao    *string `json:"dataEmissao,omitempty"`
	DataVencimento *string `json:"dataVencimento,omitempty"`
}

type RpsConfig struct {
	Serie  *string `json:"serie,omitempty"`
	Numero *int    `json:"numero,omitempty"`
	Lote   *int    `json:"lote,omitempty"`
}
