package plugnotas

import (
	"fmt"
	"strconv"
	"time"
)

const (
	ConfirmacaoDaOperacao     = "CONFIRMACAODAOPERACAO"
	DesconhecimentoDaOperacao = "DESCONHECIMENTODAOPERACAO"
	OperacaoNaoRealizada      = "OPERACAONAOREALIZADA"
)

type DFeConfig struct {
	Ativo *bool `json:"ativo,omitempty"`
}

type DFe struct {
	Id                                    *string             `json:"id,omitempty"`
	IdIntegracao                          *string             `json:"idIntegracao,omitempty"`
	Codigo                                *string             `json:"codigo,omitempty"`
	Numero                                *string             `json:"numero,omitempty"`
	Serie                                 *string             `json:"serie,omitempty"`
	Chave                                 *string             `json:"chave,omitempty"`
	Protocol                              *string             `json:"protocol,omitempty"`
	Recibo                                *string             `json:"recibo,omitempty"`
	Status                                *string             `json:"status,omitempty"`
	Finalidade                            *string             `json:"finalidade,omitempty"`
	Natureza                              *string             `json:"natureza,omitempty"`
	DataEmissao                           *string             `json:"dataEmissao,omitempty"`
	DataSaidaEntrada                      *string             `json:"dataSaidaEntrada,omitempty"`
	DataAutorizacao                       *string             `json:"dataAutorizacao,omitempty"`
	TipoImpressao                         *string             `json:"tipoImpressao,omitempty"`
	TipoEmissao                           *string             `json:"tipoEmissao,omitempty"`
	CodigoIdentificacaoDestino            *string             `json:"codigoIdentificacaoDestino,omitempty"`
	CodigoMunicipioFatoGerador            *string             `json:"codigoMunicipioFatoGerador,omitempty"`
	InformacoesComplementaresContribuinte *string             `json:"informacoesComplementaresContribuinte,omitempty"`
	InformacoesComplementares             *string             `json:"informacoesComplementares,omitempty"`
	VersaoManual                          *string             `json:"versaoManual,omitempty"`
	PDF                                   *string             `json:"pdf,omitempty"`
	XML                                   *string             `json:"xml,omitempty"`
	Intermediador                         *int                `json:"intermediador,omitempty"`
	Producao                              *bool               `json:"producao,omitempty"`
	Saida                                 *bool               `json:"saida,omitempty"`
	Presencial                            *bool               `json:"presencial,omitempty"`
	ConsumidorFinal                       *bool               `json:"consumidorFinal,omitempty"`
	Enviaremail                           *bool               `json:"enviaremail,omitempty"`
	Retorno                               interface{}         `json:"retorno,omitempty"`
	NotaReferenciada                      *NotasReferenciadas `json:"notaReferenciada,omitempty"`
	Emitente                              *Emitente           `json:"emitente,omitempty"`
	Destinatario                          *Destinatario       `json:"destinatario,omitempty"`
	Itens                                 []ItemDFe           `json:"itens,omitempty"`
	Total                                 *Total              `json:"total,omitempty"`
	Transporte                            *TransporteDFe      `json:"transporte,omitempty"`
	Pagamentos                            []Pagamento         `json:"pagamentos,omitempty"`
	Cobranca                              *Cobranca           `json:"cobranca,omitempty"`
	LocalEntrega                          *Local              `json:"localEntrega,omitempty"`
	LocalRetirada                         *Local              `json:"localRetirada,omitempty"`
	ResponsavelAutorizado                 []Responsavel       `json:"responsavelAutorizado,omitempty"`
	Exportacao                            *Exportacao         `json:"exportacao,omitempty"`
	Compra                                *Compra             `json:"compra,omitempty"`
	IntermediadorTransacao                *Intermediador      `json:"intermediadorTransacao,omitempty"`
	Cancelamento                          *Cancelamento       `json:"cancelamento,omitempty"`
	Observacoes                           *Observacoes        `json:"observacoes,omitempty"`
	CamposExtras                          *CamposExtras       `json:"camposExtras,omitempty"`
	CreatedAt                             *time.Time          `json:"createdAt,omitempty"`
	UpdatedAt                             *time.Time          `json:"updatedAt,omitempty"`
}

type DFeListagem struct {
	Notas             []DFe  `json:"notas,omitempty"`
	HashProximaPagina string `json:"hashProximaPagina,omitempty"`
}

func (c *Client) ListarDFes(cpfCnpj string, manifestada bool, dataInicial, dataFinal *time.Time, hashProximaPagina *string) (list DFeListagem, err error) {
	params := Params{
		"cpfCnpj":     cpfCnpj,
		"manifestada": strconv.FormatBool(manifestada),
	}
	if dataInicial != nil {
		params["dataInicial"] = dataInicial.Format("2006-01-02")
	}
	if dataFinal != nil {
		params["dataFinal"] = dataFinal.Format("2006-01-02")
	}
	if hashProximaPagina != nil {
		params["hashProximaPagina"] = hashProximaPagina
	}
	err = c.Get("nfe/destinada", params, nil, &list)
	return
}

func (c *Client) ManifestarNota(idNota, operacao, justificativa string) error {
	params := Params{"operacao": operacao, "justificativa": justificativa}
	return c.Get(fmt.Sprintf("nfe/%s/manifestacao", idNota), params, nil, nil)
}
