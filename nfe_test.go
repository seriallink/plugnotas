package plugnotas

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmitirNFe(t *testing.T) {
	nfe := []NFe{{
		Natureza:        StringPtr(""),
		Presencial:      BoolStr(true),
		ConsumidorFinal: BoolStr(true),
		Emitente: &Emitente{
			CpfCnpj: StringPtr(""),
		},
		Destinatario: &Destinatario{
			CpfCnpj:            StringPtr(""),
			RazaoSocial:        StringPtr(""),
			InscricaoMunicipal: StringPtr(""),
			Endereco: &Endereco{
				TipoLogradouro:  StringPtr(""),
				Logradouro:      StringPtr(""),
				Numero:          StringPtr(""),
				Complemento:     StringPtr(""),
				TipoBairro:      StringPtr(""),
				Bairro:          StringPtr(""),
				CodigoPais:      StringPtr(""),
				DescricaoPais:   StringPtr(""),
				CodigoCidade:    StringPtr(""),
				DescricaoCidade: StringPtr(""),
				Estado:          StringPtr(""),
				Cep:             StringPtr(""),
			},
		},
		Transporte: &Transporte{
			Transportador: &Transportador{
				Nome: StringPtr(""),
				Cpf:  StringPtr(""),
			},
			Veiculo: &Veiculo{
				Placa: StringPtr(""),
				UF:    StringPtr(""),
			},
			Volumes: &Volumes{
				Quantidade: FloatPtr(0),
				Especie:    StringPtr(""),
			},
		},
		Itens: []Item{{
			Descricao: StringPtr(""),
			NCM:       StringPtr(""),
			CEST:      StringPtr(""),
			CFOP:      StringPtr(""),
			Unidade: &ComercialTributavelString{
				Comercial:  StringPtr(""),
				Tributavel: StringPtr(""),
			},
			Quatidade: &ComercialTributavelNumber{
				Comercial:  FloatPtr(0),
				Tributavel: FloatPtr(0),
			},
			ValorUnitario: &ComercialTributavelNumber{
				Comercial:  FloatPtr(0),
				Tributavel: FloatPtr(0),
			},
			Valor: FloatPtr(1),
			Tributos: &Tributos{
				ICMS: &ICMS{
					Origem: StringPtr("0"),
					CST:    StringPtr("00"),
					BaseCalculo: &BaseCalculoIcms{
						ModalidadeDeterminacao: FloatPtr(0),
						Valor:                  FloatPtr(0),
					},
					Aliquota: FloatPtr(0),
					Valor:    FloatPtr(0),
				},
				PIS: &PIS{
					CST:      StringPtr("00"),
					Aliquota: FloatPtr(0),
					Valor:    FloatPtr(0),
				},
				COFINS: &COFINS{
					CST:      StringPtr("00"),
					Aliquota: FloatPtr(0),
					Valor:    FloatPtr(0),
				},
			},
		}},
		Pagamentos: []Pagamento{{
			AVista: BoolStr(true),
			Valor:  FloatPtr(1),
		}},
		InformacoesComplementares: StringPtr(""),
	}}
	resp, err := client.EmitirNFe(nfe)
	assert.NoError(t, err)
	spew.Dump(resp, err)
}

func TestConsultarNFe(t *testing.T) {
	response, err := client.ConsultarNFe("")
	assert.NoError(t, err)
	spew.Dump(response)
}

func TestCancelarNFe(t *testing.T) {
	response, err := client.CancelarNFe("", "")
	assert.NoError(t, err)
	spew.Dump(response)
}
