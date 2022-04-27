package plugnotas

type EnderecoTransportador struct {
	Logradouro      *string `json:"logradouro,omitempty"`
	DescricaoCidade *string `json:"descricaoCidade,omitempty"`
	UF              *string `json:"uf,omitempty"`
}
