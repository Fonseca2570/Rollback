/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a server for Greeter service.
package main

import (
	"context"
	"log"
	"net"
	"service_one/database"
	"service_one/initialize"

	"google.golang.org/grpc"
	pb "service_one/proto"
)

const (
	port = ":50052"
)

// server is used to implement proto.GreeterServer.
type server struct {
	pb.UnimplementedServiceServer
}

func (s *server) UpdateLastName(ctx context.Context, in *pb.UpdateLastNameRequest) (*pb.UpdateLastNameResponse, error) {
	response := &pb.UpdateLastNameResponse{}

	_, err := database.DbConnection.Update("user_one").
		Set("last_name", in.LastName).
		Where("id_users = ?", in.UsersId).
		Exec()

	if err != nil {
		response.Error = err.Error()
	}

	response.TableName = "user_two"
	response.RowId = in.UsersId

	return response, nil
}

func main() {
	initialize.Initialize()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
