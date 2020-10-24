package net

type NetKind string

const (
	NetKindTCP NetKind = "tcp"
)

var (
	_ transport.Transporter = (*Transport)(nil)
)

type Transport struct {
	endpoint      string
	operation     string
}

func NewTransport(endpoint string,
	operation string,
) *Transport {
	return &Transport{
		endpoint:  endpoint,
		operation: operation,
	}
}

func (tr *Transport) Kind() transport.Kind {
	return transport.KindTCP
}

func (tr *Transport) Endpoint() string {
	return tr.endpoint
}

func (tr *Transport) Operation() string {
	return tr.operation
}
