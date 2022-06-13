package app

import (
	"context"
	"crypto/tls"
	"net/http"
	"time"

	"github.com/opensearch-project/opensearch-go"
	"github.com/opensearch-project/opensearch-go/opensearchapi"
)

func newOpensearchClient(c OpensearchConfig) (*opensearch.Client, error) {
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.TLSClientConfig = tlsConfig

	return opensearch.NewClient(opensearch.Config{
		Addresses: c.Addresses,
		Username:  c.Username,
		Password:  c.Password,
		Transport: transport,
	})
}

func pingOpenSearch(ctx context.Context, client *opensearch.Client) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	pingRequest := opensearchapi.PingRequest{}
	_, err := pingRequest.Do(ctx, client)
	return err
}
