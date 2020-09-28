package plugnotas

type Importacao struct {
	Numero                        *string            `json:"numero,omitempty"`
	DataEmissao                   *string            `json:"dataEmissao,omitempty"`
	ViaTransporte                 *string            `json:"viaTransporte,omitempty"`
	ValorAfrmm                    *string            `json:"valorAfrmm,omitempty"`
	FormaImportacao               *string            `json:"formaImportacao,omitempty"`
	CodigoExportador              *string            `json:"codigoExportador,omitempty"`
	PercentualMercadoriaDevolvida *float64           `json:"percentualMercadoriaDevolvida,omitempty"`
	ValorIpiDevolvido             *float64           `json:"valorIpiDevolvido,omitempty"`
	ValorAproximadoTributos       *float64           `json:"valorAproximadoTributos,omitempty"`
	Desembaraco                   *Desembaraco       `json:"desembaraco,omitempty"`
	Adquirente                    *Adquirente        `json:"adquirente,omitempty"`
	Adicoes                       *Adicoes           `json:"adicoes,omitempty"`
	ImpostoImportacao             *ImpostoImportacao `json:"impostoImportacao,omitempty"`
}
