package http

import (
	"context"
	"io"
	"net/http"
	"time"

	. "github.com/iggy-rs/iggy-go-client/contracts"
	"github.com/pkg/errors"
)

var (
	// Multiplex the client transport
	muxClient *http.Client
)

type IggyHttpClient struct {
	client *http.Client
}

func init() {
	dt, _ := http.DefaultTransport.(*http.Transport)
	t := dt.Clone()

	// http client tunning from
	// https://www.loginradius.com/blog/engineering/tune-the-go-http-client-for-high-performance/
	t.MaxIdleConns = 100
	t.MaxConnsPerHost = 100
	t.MaxIdleConnsPerHost = 100
	client := &http.Client{
		Timeout:   10 * time.Second,
		Transport: t,
	}
	muxClient = client
}

func NewHttpMessageStream(url string) (*IggyHttpClient, error) {
	return &IggyHttpClient{client: muxClient}, nil
}

func (tms *IggyHttpClient) sendAndFetchResponse(ctx context.Context, uri string, method string, payLoad io.Reader, contentType string) ([]byte, int, error) {

	req, err := http.NewRequest(method, uri, payLoad)
	if err != nil {
		return nil, 0, err
	}

	req = req.WithContext(ctx)

	if contentType != "" {
		req.Header.Add("Content-Type", contentType)
	}

	resp, err := tms.client.Do(req)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	defer resp.Body.Close()

	buffer, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, errors.New("read response body failed")
	}

	return buffer, resp.StatusCode, nil
}
