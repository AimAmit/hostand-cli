package api

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/aimamit/hostand-cli/proto"
)

func BuildImage(domain, version string, buffer bytes.Buffer) error {
	stream, err := GrpcClient.Docker.FileUpload(context.Background())
	if err != nil {
		return err
	}
	buf := make([]byte, 1024)

	err = stream.Send(&proto.FileRequest{
		Data: &proto.FileRequest_AppVersion{
			AppVersion: &proto.AppVersion{
				Domain:  domain,
				Version: version,
			},
		},
	})
	if err != nil {
		fmt.Println("cannot send chunk to server: ", err, stream.RecvMsg(nil))
		return err
	}

	for {
		n, err := buffer.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		req := &proto.FileRequest{
			Data: &proto.FileRequest_Chunk{
				Chunk: buf[:n],
			},
		}

		err = stream.Send(req)
		if err != nil {
			fmt.Println("cannot send chunk to server: ", err, stream.RecvMsg(nil))
			return err
		}
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}
