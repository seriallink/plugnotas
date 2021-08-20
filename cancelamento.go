package plugnotas

type Cancelamento struct {
	XML   interface{} `json:"xml,omitempty"`
	Error interface{} `json:"error,omitempty"`
}
