package plugnotas

type Tributos struct {
	Partilha *Partilha `json:"partilha,omitempty"`
	PIS      *PIS      `json:"pis,omitempty"`
	COFINS   *COFINS   `json:"cofins,omitempty"`
	IPI      *IPI      `json:"ipi,omitempty"`
	ISSQN    *ISSQN    `json:"issqn,omitempty"`
	ICMS     *ICMS     `json:"icms,omitempty"`
}
