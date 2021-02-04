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

func (c *Client) ListarDFes(cpfCnpj string, manifestada bool, dataInicial, dataFinal *time.Time) (list NFeListagemCompleta, err error) {
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
	err = c.Get("nfe/destinada", params, nil, &list)
	return
}

func (c *Client) ManifestarNota(idNota, operacao, justificativa string) error {
	params := Params{"operacao": operacao, "justificativa": justificativa}
	return c.Get(fmt.Sprintf("nfe/%s/manifestacao", idNota), params, nil, nil)
}
