package plugnotas

type Cancelamento struct {
	XML   interface{} `json:"xml,omitempty"`
	Error interface{} `json:"error,omitempty"`
}

type CancelamentoStatus struct {
	Message *string           `json:"message,omitempty"`
	Data    *CancelamentoData `json:"data,omitempty"`
}

type CancelamentoData struct {
	XML           *string `json:"xml,omitempty"`
	Justificativa *string `json:"justificativa,omitempty"`
	CStat         *int    `json:"cStat,omitempty"`
	Date          *string `json:"date,omitempty"`
	Protocolo     *string `json:"protocolo,omitempty"`
	RespostaSefaz *string `json:"respostaSefaz,omitempty"`
	Status        *string `json:"status,omitempty"`
}
