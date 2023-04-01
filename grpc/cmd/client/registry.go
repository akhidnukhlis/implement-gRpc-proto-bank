package client

import "github.com/akhidnukhlis/implement-gRpc-microservice/grpc/pb"

type AuthorClient struct {
	Author pb.AuthorServiceClient
}
