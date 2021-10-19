package transaction

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
	pb "transaction/proto"
)

const(
	address1     = "localhost:50051"
)

func (u *UpdateModel) Update() error {
	fmt.Println(u)

	connServiceOne()

	return nil
}


func connServiceOne() error {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address1, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewServiceClient(conn)

	user := &pb.UpdateFirstNameRequest{
		FirstName: "joao",
		UsersId: 1,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.UpdateFirstName(ctx, user)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.TableName)
	return nil
}