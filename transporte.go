package plugnotas

type Transporte struct {
	ModalidadeFrete *string        `json:"modalidadeFrete,omitempty"`
	Transportador   *Transportador `json:"transportador,omitempty"`
	RetencaoICMS    *RetencaoIcms  `json:"RetencaoICMS,omitempty"`
	Veiculo         *Veiculo       `json:"veiculo,omitempty"`
	Reboque         *Reboque       `json:"reboque,omitempty"`
	Volumes         *Volumes       `json:"volumes,omitempty"`
}
