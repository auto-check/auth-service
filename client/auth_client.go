package client

import (
	authpb "github.com/auto-check/protocol-buffer/golang/auth"
	"google.golang.org/grpc"
	"sync"
)

var (
	once sync.Once
	cli authpb.AuthClient
)

func GetAuthClient() authpb.AuthClient {
	once.Do(func() {
		conn, _ := grpc.Dial(
			"127.0.0.1:5000",
			grpc.WithInsecure(),
			grpc.WithBlock())

		cli = authpb.NewAuthClient(conn)
	})

	return cli
}

