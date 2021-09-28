.PHONY: protoc
GOPATH = /home/aimamit/go


protoc:
    protoc -I . ./proto/auth.proto --go_out=. --go-grpc_out=. --proto_path=.:${GOPATH}/src
