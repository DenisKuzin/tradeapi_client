package traceapi_client

import (
	"context"
	"crypto/tls"

	tradeapi_grpc "github.com/DenisKuzin/tradeapi/grpc/tradeapi/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

type Client struct {
	ctx       context.Context
	conn      *grpc.ClientConn
	portfolio tradeapi_grpc.PortfoliosClient
}

func NewClient(login string, token string, ctx context.Context) (*Client, error) {
	endpoint := "trade-api.finam.ru:443"
	tlsConfig := tls.Config{MinVersion: tls.VersionTLS13}
	conn, err := grpc.DialContext(ctx, endpoint, grpc.WithTransportCredentials(credentials.NewTLS(&tlsConfig)))
	if err != nil {
		return nil, err
	}
	ctx = metadata.AppendToOutgoingContext(ctx, "x-api-key", token)

	client := Client{
		ctx:       ctx,
		conn:      conn,
		portfolio: tradeapi_grpc.NewPortfoliosClient(conn),
	}
	return &client, nil
}
