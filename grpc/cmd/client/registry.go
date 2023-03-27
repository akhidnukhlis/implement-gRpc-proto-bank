package client

import "github.com/akhidnukhlis/implement-gRpc-microservice/grpc/pb"

type AuthorClient struct {
	User pb.AuthorServiceClient
}
