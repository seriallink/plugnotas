package plugnotas

type Integracoes struct {
	PlugStore *PlugStore `json:"plugStore,omitempty"`
}

type PlugStore struct {
	Ativo *bool   `json:"ativo,omitempty"`
	Email *string `json:"email,omitempty"`
	Senha *string `json:"senha,omitempty"`
}
