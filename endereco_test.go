package plugnotas

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConsultaEnderece(t *testing.T) {
	response, err := client.ConsultaEndereco("")
	assert.NoError(t, err)
	spew.Dump(response)
}
