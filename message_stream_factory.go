package iggy

import (
	"errors"

	"github.com/iggy-rs/iggy-go-client/net/http"
	"github.com/iggy-rs/iggy-go-client/net/tcp"

	iggcon "github.com/iggy-rs/iggy-go-client/contracts"
)

type IMessageStreamFactory interface {
	CreateStream(config iggcon.IggyConfiguration) (MessageStream, error)
}

type IggyClientFactory struct{}

func (msf *IggyClientFactory) CreateMessageStream(config iggcon.IggyConfiguration) (MessageStream, error) {
	if config.Protocol == iggcon.Tcp {
		tcpMessageStream, err := tcp.NewTcpMessageStream(config.BaseAddress)
		if err != nil {
			return nil, err
		}
		return tcpMessageStream, nil
	}

	if config.Protocol == iggcon.Http {
		httpMessageStream, err := http.NewHttppMessageStream(config.BaseAddress)
		if err != nil {
			return nil, err
		}
		return httpMessageStream, nil
	}

	return nil, errors.New("unsupported protocol")
}
