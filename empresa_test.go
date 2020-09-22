package plugnotas

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListaEmpresas(t *testing.T) {
	response, err := client.ListaEmpresas()
	assert.NoError(t, err)
	spew.Dump(response)
}

func TestConsultaEmpresa(t *testing.T) {
	response, err := client.ConsultaEmpresa("")
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
