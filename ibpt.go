package plugnotas

type IBPT struct {
	Simplificado *Simplificado `json:"simplificado,omitempty"`
	Detalhado    *Detalhado    `json:"detalhado,omitempty"`
}
