package grpc

import (
	context "context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	srv *grpc.Server
	UnimplementedImageUploaderServer
}

func NewServer() *Server {
	return &Server{
		srv: grpc.NewServer(),
	}
}

func (s *Server) Start(port string) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))

	if err != nil {
		return err
	}

	log.Printf("server listening at %v", lis.Addr())

	return s.srv.Serve(lis)
}

func (s *Server) Upload(ctx context.Context, in *Image) (*ImageUploadReply, error) {
	return &ImageUploadReply{Location: "hello world"}, nil
}
