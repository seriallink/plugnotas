package plugnotas

type Partilha struct {
	UfDestino   *UfDestino   `json:"ufDestino,omitempty"`
	UfRemetente *UfRemetente `json:"ufRemetente,omitempty"`
}
