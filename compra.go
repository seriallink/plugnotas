package plugnotas

type Compra struct {
	NumeroCompra *string `json:"numeroCompra,omitempty"`
	PedidoCompra *string `json:"pedidoCompra,omitempty"`
}
