package plugclient

import (
	"github.com/fxamacker/cbor/v2"
	"github.com/mjwhodur/plugkit/client"
)

type MisoPlugClient struct {
}

func (m MisoPlugClient) Handle(kind string, payload *cbor.RawMessage) {
	//TODO implement me
	panic("implement me")
}

func (m MisoPlugClient) Mount(c *client.RawStreamClient) {
	//TODO implement me
	panic("implement me")
}

func (m MisoPlugClient) CloseSignal() {
	//TODO implement me
	panic("implement me")
}
