package transaction

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
	pb "transaction/proto"
)

const (
	address1 = "localhost:50051"
	address2 = "localhost:50052"
	address3 = "localhost:50053"
)

func (u *UpdateModel) Update() error {
	fmt.Println(u)

	connServiceOne(u)

	return nil
}

func connServiceOne(u *UpdateModel) error {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address1, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewServiceClient(conn)

	user := &pb.UpdateFirstNameRequest{
		FirstName: u.FirstName,
		UsersId:   int64(u.IdUser),
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
