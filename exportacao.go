package plugnotas

type Exportacao struct {
	EstadoEmbarque         string `json:"estadoEmbarque,omitempty"`
	DescricaoLocalEmbarque string `json:"descricaoLocalEmbarque,omitempty"`
	DescricaoLocalDespacho string `json:"descricaoLocalDespacho,omitempty"`
}
