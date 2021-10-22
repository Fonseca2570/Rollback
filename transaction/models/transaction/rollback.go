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
	GetError() string
}

func (u *UpdateResponseServiceOne) Rollback() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address1, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewServiceClient(conn)

	rollback := &pb.RollbackRequest{
		TableName: u.TableName,
		Id:        u.RowId,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Rollback(ctx, rollback)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
}

func (u *UpdateResponseServiceOne) GetError() string{
	return u.Error
}
