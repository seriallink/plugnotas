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
		CpfCnpj:                  "",
		RazaoSocial:              "",
		NomeFantasia:             "",
		InscricaoEstadual:        "",
		Email:                    "",
		SimplesNacional:          true,
		RegimeTributario:         1,
		RegimeTributarioEspecial: 5,
		Endereco: &Endereco{
			TipoLogradouro: "",
			Logradouro:     "",
			Numero:         "",
			Complemento:    "",
			//TipoBairro:      "",
			Bairro:     "",
			CodigoPais: "",
			//DescricaoPais:   "",
			CodigoCidade:    "",
			DescricaoCidade: "",
			Estado:          "",
			Cep:             "",
		},
	}
	err := client.CadastrarEmpresa(empresa)
	assert.NoError(t, err)
	//assert.Equal(t, "", response.Data.Id)
	//spew.Dump(err, response.Data)
}

func TestAlterarEmpresa(t *testing.T) {

	empresa := &Empresa{
		CpfCnpj:     "",
		Certificado: "",
	}

	err := client.AlterarEmpresa(empresa)
	assert.NoError(t, err)

}
