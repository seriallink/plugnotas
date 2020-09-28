package plugnotas

type Adquirente struct {
	Cnpj   *string `json:"cnpj,omitempty"`
	Estado *string `json:"estado,omitempty"`
}
