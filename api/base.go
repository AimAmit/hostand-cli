package api

import (
	"flag"
	"fmt"
	"net/url"

	"github.com/aimamit/hostand-cli/proto"
	"google.golang.org/grpc"
)

type Client struct {
	Docker proto.DockerServiceClient
	Auth   proto.AuthServiceClient
}

var GrpcClient Client

func Init() error {

	DockerHost := "65.0.138.22:80"
	AuthHost, _ := url.Parse("http://aimamit.stage.toppr.io:50051")
	serverAddress := flag.String("address", DockerHost, "tcp")
	flag.Parse()

	conn, err := grpc.Dial(*serverAddress, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("error connecting docker remote server: %v", err)
	}

	authConn, err := grpc.Dial(AuthHost.Host, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("error connecting auth remote server: %v", err)
	}
	GrpcClient.Docker = proto.NewDockerServiceClient(conn)
	GrpcClient.Auth = proto.NewAuthServiceClient(authConn)

	return nil
}
