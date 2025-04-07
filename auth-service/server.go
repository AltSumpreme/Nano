package auth_service

import (
	"context"
	"fmt"
	"net"

	"github.com/AltSumpreme/Nano.git/auth-service/pb/github.com/AltSumpreme/Nano/account/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	service Service
}

func listenGrpc(service Service, port int) error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	server := grpc.NewServer()
	pb.RegisterAccountServiceServer(server, &grpcServer{service})
	reflection.Register(server)
	return server.Serve(listener)
}

func (s *grpcServer) PostAccount(ctx context.Context, r *pb.PostAccountRequest) (*pb.PostAccountResponse, error) {
	account, err := s.service.PostAccount(ctx, r.Name)
	if err != nil {
		return nil, err
	}
	return &pb.PostAccountResponse{Account: &pb.Account{
		Id:   account.ID,
		Name: account.Name,
	}}, nil

}

func (s *grpcServer) GetAccount(ctx context.Context, r *pb.GetAccountRequest) (*pb.GetAccountResponse, error) {
	account, err := s.service.GetAccount(ctx, r.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetAccountResponse{Account: &pb.Account{
		Id:   account.ID,
		Name: account.Name,
	}}, nil
}

func (s *grpcServer) GetAccounts(ctx context.Context, r *pb.GetAccountsRequest) (*pb.GetAccountsResponse, error) {
	response, err := s.service.GetAccounts(ctx, r.Skip, r.Take)
	if err != nil {
		return nil, err

	}
	accounts := []*pb.Account{}
	for _, p := range response {
		accounts = append(accounts, &pb.Account{Id: p.ID, Name: p.Name})
	}
	return &pb.GetAccountsResponse{Accounts: accounts}, nil
}
