package main

import (
    "context"
    pb "github.com/falcondb/FalconLib/golib/grpc/protob"
    "google.golang.org/grpc"
    "io"
    "log"
    "time"
)

type client struct{}

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

    stream, _ := c.ServerStreamingEcho(ctx, &pb.EchoRequest{Message: name})

    for {
        greetings, err := stream.Recv()
        if err == io.EOF {
            break
        }
        if err != nil {
            log.Fatal(err)
        }
        log.Println(greetings)
    }
}
