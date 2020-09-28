package plugnotas

type Servico struct {
	IdIntegracao              *string   `json:"idIntegracao,omitempty"`
	Codigo                    *string   `json:"codigo,omitempty"`
	Discriminacao             *string   `json:"discriminacao,omitempty"`
	CodigoTributacao          *string   `json:"codigoTributacao,omitempty"`
	CNAE                      *string   `json:"cnae,omitempty"`
	CodigoCidadeIncidencia    *string   `json:"codigoCidadeIncidencia,omitempty"`
	DescricaoCidadeIncidencia *string   `json:"descricaoCidadeIncidencia,omitempty"`
	Unidade                   *string   `json:"unidade,omitempty"`
	Quantidade                *float64  `json:"quantidade,omitempty"`
	InformacoesLegais         *string   `json:"informacoesLegais,omitempty"`
	DescricaoLC116            *string   `json:"descricaoLC116,omitempty"`
	Tributavel                *bool     `json:"tributavel,omitempty"`
	ResponsavelRetencao       *string   `json:"responsavelRetencao,omitempty"`
	ISS                       *ISS      `json:"iss,omitempty"`
	Obra                      *Obra     `json:"obra,omitempty"`
	Evento                    *Evento   `json:"evento,omitempty"`
	Valor                     *Valor    `json:"valor,omitempty"`
	Deducao                   *Deducao  `json:"deducao,omitempty"`
	Retencao                  *Retencao `json:"retencao,omitempty"`
	IBPT                      *IBPT     `json:"ibpt,omitempty"`
}
