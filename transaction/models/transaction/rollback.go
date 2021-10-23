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

	_, err = c.Rollback(ctx, rollback)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
}


func (u *UpdateResponseServiceOne) GetError() string{
	return u.Error
}

func (u *UpdateResponseServiceTwo) GetError() string{
	return u.Error
}

func (u *UpdateResponseServiceTwo) Rollback() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address2, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewServiceTwoClient(conn)

	rollback := &pb.RollbackTwoRequest{
		TableName: u.TableName,
		Id:        u.RowId,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err = c.RollbackTwo(ctx, rollback)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
}


func (u *UpdateResponseServiceThree) Rollback() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address3, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewServiceThreeClient(conn)

	rollback := &pb.RollbackThreeRequest{
		TableName: u.TableName,
		Id:        u.RowId,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err = c.RollbackThree(ctx, rollback)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
}


func (u *UpdateResponseServiceThree) GetError() string{
	return u.Error
}