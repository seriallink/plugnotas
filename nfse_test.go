package plugnotas

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmitirNFSe(t *testing.T) {
	nfse := []NFSe{{
		Prestador: &Prestador{
			CpfCnpj: StringPtr(""),
		},
		Tomador: &Tomador{
			CpfCnpj:            StringPtr(""),
			RazaoSocial:        StringPtr(""),
			InscricaoMunicipal: StringPtr(""),
			Email:              StringPtr(""),
		},
		Servico: []Servico{{
			Codigo:           StringPtr(""),
			CodigoTributacao: StringPtr(""),
			Discriminacao:    StringPtr(""),
			CNAE:             StringPtr(""),
			ISS: &ISS{
				TipoTributacao: IntPtr(0),
				Exigibilidade:  IntPtr(0),
				Aliquota:       FloatPtr(0),
			},
			Valor: &Valor{
				Servico:                FloatPtr(0),
				DescontoCondicionado:   FloatPtr(0),
				DescontoIncondicionado: FloatPtr(0),
			},
		}},
	}}
	resp, err := client.EmitirNFSe(nfse)
	assert.NoError(t, err)
	spew.Dump(resp, err)
}

func TestConsultarNFSe(t *testing.T) {
	response, err := client.ConsultarNFSe("")
	assert.NoError(t, err)
	spew.Dump(response)
}

func TestDownloadPDF(t *testing.T) {
	pdf, err := client.DownloadNFSePDF("")
	assert.NoError(t, err)
	assert.NotZero(t, len(pdf))
}
