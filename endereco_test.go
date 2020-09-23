package plugnotas

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConsultarEndereco(t *testing.T) {
	response, err := client.ConsultarEndereco("")
	assert.NoError(t, err)
	spew.Dump(response)
}
