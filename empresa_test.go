package plugnotas

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListarEmpresas(t *testing.T) {
	response, err := client.ListarEmpresas()
	assert.NoError(t, err)
	spew.Dump(response)
}

func TestConsultarEmpresa(t *testing.T) {
	response, err := client.ConsultarEmpresa("")
	assert.NoError(t, err)
	spew.Dump(response)
}

func TestCadastrarEmpresa(t *testing.T) {
	empresa := &Empresa{
		CpfCnpj:                  StringPtr(""),
		RazaoSocial:              StringPtr(""),
		NomeFantasia:             StringPtr(""),
		InscricaoEstadual:        StringPtr(""),
		Email:                    StringPtr(""),
		SimplesNacional:          BoolStr(true),
		RegimeTributario:         IntPtr(1),
		RegimeTributarioEspecial: IntPtr(5),
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
	}
	err := client.CadastrarEmpresa(empresa)
	assert.NoError(t, err)
}

func TestAlterarEmpresa(t *testing.T) {

	empresa := &Empresa{
		CpfCnpj:     StringPtr(""),
		Certificado: StringPtr(""),
	}

	err := client.AlterarEmpresa(empresa)
	assert.NoError(t, err)

}
