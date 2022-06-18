package grpc

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/0xdod/imageuploadservice/internal/imageuploader"
	"google.golang.org/grpc"
)

var uploader imageuploader.ImageUploader

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
	uploader = imageuploader.NewS3Uploader()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))

	if err != nil {
		return err
	}

	// Register server method (actions the server will do)
	RegisterImageUploaderServer(s.srv, &Server{})

	log.Printf("server listening at %v", lis.Addr())

	return s.srv.Serve(lis)
}

func (s *Server) Upload(ctx context.Context, in *Image) (*ImageUploadReply, error) {
	loc, err := uploader.Upload(ctx, in.Body, in.Name)

	if err != nil {
		return nil, err
	}

	return &ImageUploadReply{Location: loc}, nil
}
