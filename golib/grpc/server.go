package sc

import (
    "context"
    "fmt"
    pb "github.com/falcondb/FalconLib/golib/grpc/protob"
    "google.golang.org/grpc"
    "log"
    "net"
)

type server struct {
     who string
}


func (*server) UnaryEcho(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
    return &pb.EchoResponse{Message: server.who + " is doing good!"}, nil
}

// ServerStreamingEcho is server side streaming.
func (*server) ServerStreamingEcho(ctx *pb.EchoRequest, st pb.Echo_ServerStreamingEchoServer) error {
    return nil
}
// ClientStreamingEcho is client side streaming.
func (*server)  ClientStreamingEcho(st pb.Echo_ClientStreamingEchoServer) error {
    return nil
}
// BidirectionalStreamingEcho is bidi streaming.
func (*server) BidirectionalStreamingEcho(st pb.Echo_BidirectionalStreamingEchoServer) error {
    return nil
}

func main() {
    lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 11123))
    if err != nil {
            panic(err)
    }

    s := grpc.NewServer()
    pb.RegisterEchoServer(s, &server{})
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }

}