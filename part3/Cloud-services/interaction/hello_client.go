package main

import (
	pb "./hello"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func main() {
	address:="localhost:55555"
	conn,err:=grpc.Dial(address,grpc.WithInsecure())
	if err != nil {
		return
	}
	defer conn.Close()
	c:=pb.NewHelloClient(conn)
	name:="Inigo Montoya"
	hr:=&pb.HelloRequest{Name: name}
	r,err:=c.Say(context.Background(),hr)
	if err != nil {
		return
	}
	fmt.Println(r.Message)
}