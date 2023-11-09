package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

type Client struct {
	ctx  context.Context
	conn *grpc.ClientConn
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
		ctx:  ctx,
		conn: conn,
	}
	return &client, nil
}

func (f *Client) CloseConnection() {
	f.conn.Close()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Begin read finam login")
	login_bytes, err := os.ReadFile("/home/deniskuzin/.finam/login")
	if err != nil {
		log.Panicln(err)
	}
	login := string(login_bytes)
	fmt.Println(login)
	fmt.Println("End read finam login")
	fmt.Println("Begin read finam token")
	token_bytes, err := os.ReadFile("/home/deniskuzin/.finam/token")
	if err != nil {
		log.Panicln(err)
	}
	token := string(token_bytes)
	fmt.Println("End read finam token")
	ctx := context.Background()
	client, err := NewClient(string(login), string(token), ctx)
	if err != nil {
		log.Panicln(err)
	}
	defer client.CloseConnection()
	// res, err := client.portfolio.GetPortfolio(true, true, true, true)
	// if err != nil {
	// 	log.Panicln(err)
	// }
	// log.Printf("Входящая оценка портфеля в рублях: %f", res.Balance)
	// log.Println(res.Positions)
}
