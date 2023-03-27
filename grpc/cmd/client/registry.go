package client

import "implement-gRpc-microservice/grpc/pb"

type AuthorClient struct {
	User pb.AuthorServiceClient
}
