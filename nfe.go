package plugnotas

import "fmt"

type NFe struct {
	Codigo                     *string           `json:"codigo,omitempty"`
	Numero                     *string           `json:"numero,omitempty"`
	Serie                      *string           `json:"serie,omitempty"`
	Chave                      *string           `json:"chave,omitempty"`
	Recibo                     *string           `json:"recibo,omitempty"`
	Status                     *string           `json:"status,omitempty"`
	Finalidade                 *string           `json:"finalidade,omitempty"`
	Natureza                   *string           `json:"natureza,omitempty"`
	VersaoManual               *string           `json:"versaoManual,omitempty"`
	DataEmissao                *string           `json:"dataEmissao,omitempty"`
	DataSaidaEntrada           *string           `json:"dataSaidaEntrada,omitempty"`
	DataAutorizacao            *string           `json:"dataAutorizacao,omitempty"`
	TipoImpressao              *string           `json:"tipoImpressao,omitempty"`
	TipoEmissao                *string           `json:"tipoEmissao,omitempty"`
	CodigoIdentificacaoDestino *string           `json:"codigoIdentificacaoDestino,omitempty"`
	CodigoMunicipioFatoGerador *string           `json:"codigoMunicipioFatoGerador,omitempty"`
	InformacoesComplementares  *string           `json:"informacoesComplementares,omitempty"`
	Saida                      *bool             `json:"saida,omitempty"`
	Presencial                 *bool             `json:"presencial,omitempty"`
	ConsumidorFinal            *bool             `json:"consumidorFinal,omitempty"`
	Enviaremail                *bool             `json:"enviaremail,omitempty"`
	NotaReferenciada           *NotaReferenciada `json:"notaReferenciada,omitempty"`
	Emitente                   *Emitente         `json:"emitente,omitempty"`
	Destinatario               *Destinatario     `json:"destinatario,omitempty"`
	Total                      *Total            `json:"total,omitempty"`
	Transporte                 *Transporte       `json:"transporte,omitempty"`
	Cobranca                   *Cobranca         `json:"cobranca,omitempty"`
	Itens                      []Item            `json:"itens,omitempty"`
	Pagamentos                 []Pagamento       `json:"pagamentos,omitempty"`
}

type NFeResumo struct {
	Id               string  `json:"id"`
	Codigo           string  `json:"codigo"`
	Emissao          string  `json:"emissao"`
	Status           string  `json:"status"`
	Destinada        bool    `json:"destinada"`
	Emitente         string  `json:"emitente"`
	Destinatario     string  `json:"destinatario"`
	Valor            float64 `json:"valor"`
	Numero           string  `json:"numero"`
	Serie            string  `json:"serie"`
	Chave            string  `json:"chave"`
	Protocolo        string  `json:"protocolo"`
	DataAutorizacao  string  `json:"dataAutorizacao"`
	Mensagem         string  `json:"mensagem"`
	Pdf              string  `json:"pdf"`
	Xml              string  `json:"xml"`
	DataCancelamento string  `json:"dataCancelamento"`
	Erro             string  `json:"erro"`
}

type NFeObject struct {
	Ativo        bool       `json:"ativo,omitempty"`
	TipoContrato int        `json:"tipoContrato,omitempty"`
	Config       *NFeConfig `json:"config,omitempty"`
}

type NFeConfig struct {
	Producao          bool `json:"producao,omitempty"`
	ImpressaoFcp      bool `json:"impressaoFcp,omitempty"`
	ImpressaoPartilha bool `json:"impressaoPartilha,omitempty"`
	Serie             int  `json:"serie,omitempty"`
	Numero            int  `json:"numero,omitempty"`
}

func (c *Client) EmitirNFe(nfes []NFe) (*NfResponse, error) {
	resp := new(NfResponse)
	err := c.Post("nfe/", nfes, nil, resp)
	return resp, err
}

func (c *Client) ConsultarNFe(id string) (*NFeResumo, error) {
	var nfes []NFeResumo
	err := c.Get(fmt.Sprintf("nfe/%s/resumo", id), nil, nil, &nfes)
	if len(nfes) == 1 {
		return &nfes[0], err
	}
	return &NFeResumo{}, err
}

func (c *Client) CancelarNFe(id, justificativa string) (*NfResponse, error) {
	params := map[string]string{"justificativa": "Erro na emissão da Nota Fiscal eletrônica."}
	if justificativa != "" {
		params["justificativa"] = justificativa
	}
	resp := new(NfResponse)
	err := c.Post(fmt.Sprintf("nfe/%s/cancelamento", id), params, nil, resp)
	return resp, err
}
