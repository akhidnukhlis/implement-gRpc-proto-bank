package client

import "github.com/akhidnukhlis/implement-gRpc-proto-bank/grpc/pb"

type AuthorClient struct {
	Author pb.AuthorServiceClient
}
