package transaction

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
	pb "transaction/proto"
)

type RollbackInterface interface {
	Rollback()
}


func (u *UpdateResponseServiceOne) Rollback() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address1, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewServiceClient(conn)

	user := &pb.RollbackRequest{
		FirstName: u.FirstName,
		UsersId:   int64(u.IdUser),
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Rollback(ctx, user)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.TableName)
}