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
	var results []RollbackInterface
	var errorCount int

	results = append(results, connServiceOne(u))
	results = append(results, connServiceTwo(u))
	results = append(results, connServiceThree(u))

	for _, result := range results {
		if result.GetError() != "" {
			errorCount++
		}
	}

	if errorCount > 0 {
		// Service three will crate an error so the rollback happens
		fmt.Println("Error count is bigger then 0")
		fmt.Println(errorCount)
		for _, result := range results {
			if result.GetError() == "" {
				result.Rollback()
			}
		}
	}

	return nil
}

func connServiceOne(u *UpdateModel) *UpdateResponseServiceOne {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address1, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Printf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewServiceClient(conn)

	user := &pb.UpdateFirstNameRequest{
		FirstName: u.FirstName,
		UsersId:   int64(u.IdUser),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	r, err := c.UpdateFirstName(ctx, user)
	if err != nil {
		r.Error = err.Error()
	}

	return &UpdateResponseServiceOne{
		Error:     r.Error,
		TableName: r.TableName,
		RowId:     r.RowId,
	}
}

func connServiceTwo(u *UpdateModel) *UpdateResponseServiceTwo {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address2, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Printf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewServiceTwoClient(conn)

	user := &pb.UpdateLastNameRequest{
		LastName: u.LastName,
		UsersId:  int64(u.IdUser),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	r, err := c.UpdateLastName(ctx, user)
	if err != nil {
		r.Error = err.Error()
	}

	return &UpdateResponseServiceTwo{
		Error:     r.Error,
		TableName: r.TableName,
		RowId:     r.RowId,
	}
}

func connServiceThree(u *UpdateModel) *UpdateResponseServiceThree {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address3, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewServiceThreeClient(conn)

	user := &pb.UpdateEmailRequest{
		Email:   u.Email,
		UsersId: int64(u.IdUser),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	r, err := c.UpdateEmail(ctx, user)
	if err != nil {
		r.Error = err.Error()
	}

	return &UpdateResponseServiceThree{
		Error:     r.Error,
		TableName: r.TableName,
		RowId:     r.RowId,
	}
}
