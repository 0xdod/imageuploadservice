package grpc

import (
	"context"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	conn   *grpc.ClientConn
	client ImageUploaderClient
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) DialServer(addr string) error {
	// Set up a connection to the server.
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return err
	}

	c.conn = conn
	c.client = NewImageUploaderClient(conn)

	return nil
}

func (c *Client) Close() {
	c.conn.Close()
}

func (c *Client) UploadImage(ctx context.Context, name string, data []byte) (string, error) {
	res, err := c.client.Upload(ctx, &Image{Name: name, Body: data})

	if err != nil {
		return "", err
	}

	return res.Location, nil
}
