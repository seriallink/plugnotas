package plugnotas

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListaCertificados(t *testing.T) {
	certificados, err := client.ListaCertificados()
	assert.NoError(t, err)
	spew.Dump(certificados)
}

func TestUploadCertificate(t *testing.T) {
	arquivo := ""
	senha := ""
	id, err := client.UploadCertificate(arquivo, senha, "")
	assert.NoError(t, err)
	spew.Dump(id)
}
