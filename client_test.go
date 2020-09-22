package plugnotas

var (
	client *Client
)

func init() {

	client = NewClient(
		"",
		false,
	)

}
