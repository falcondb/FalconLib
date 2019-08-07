package main

import (
    "context"
    pb "github.com/falcondb/FalconLib/golib/grpc/protob"
    "google.golang.org/grpc"
    "log"
    "time"
)

type client struct{}

func (c *client) UnaryEcho(ctx context.Context, in *pb.EchoRequest, opts ...grpc.CallOption) (*pb.EchoResponse, error) {
    return nil, nil
}

func (c *client) ServerStreamingEcho(ctx context.Context, in *pb.EchoRequest, opts ...grpc.CallOption) (pb.Echo_ServerStreamingEchoClient, error) {
    return nil, nil
}

func (c *client) ClientStreamingEcho(ctx context.Context, opts ...grpc.CallOption) (pb.Echo_ClientStreamingEchoClient, error) {
    return nil, nil
}

func (c *client) BidirectionalStreamingEcho(ctx context.Context, opts ...grpc.CallOption) (pb.Echo_BidirectionalStreamingEchoClient, error) {
    return nil, nil
}

func main() {
    conn, err := grpc.Dial("localhost:11123", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    c := pb.NewEchoClient(conn)

    name := "I am a client"

    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()
    r, err := c.UnaryEcho(ctx, &pb.EchoRequest{Message: name})
    if err != nil {
        log.Fatalf("could not greet: %v", err)
    }
    log.Printf("Greeting: %s", r.Message)
}
