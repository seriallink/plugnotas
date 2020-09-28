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
	id, err := client.CadastrarCertificado("", "", "")
	assert.NoError(t, err)
	spew.Dump(id)
}
