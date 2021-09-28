package api

import (
	"context"
	"fmt"

	"github.com/aimamit/hostand-cli/proto"
	"google.golang.org/grpc/status"
)

func Signup(email, password string) (string, error) {
	request := &proto.SignupRequest{
		Email:    email,
		Password: password,
	}
	response, err := GrpcClient.Auth.Signup(context.Background(), request)
	if err != nil {
		return "", fmt.Errorf(status.Convert(err).Message())
	}

	return response.Token, nil
}

func Signin(email, password string) (string, error) {
	request := &proto.LoginRequest{
		Email:    email,
		Password: password,
	}
	response, err := GrpcClient.Auth.Login(context.Background(), request)
	if err != nil {
		return "", fmt.Errorf(status.Convert(err).Message())
	}

	return response.Token, nil
}
