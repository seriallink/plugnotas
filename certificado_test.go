package plugnotas

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListaCertificados(t *testing.T) {
	certificados, err := client.ListarCertificados()
	assert.NoError(t, err)
	spew.Dump(certificados)
}

func TestUploadCertificate(t *testing.T) {
	arquivo := ""
	senha := ""
	id, err := client.CadastrarCertificado(arquivo, senha, "")
	assert.NoError(t, err)
	spew.Dump(id)
}
