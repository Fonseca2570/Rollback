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
	"fmt"
	"github.com/mailru/dbr"
	"log"
	"net"
	"service_one/database"
	"service_one/initialize"

	"google.golang.org/grpc"
	pb "service_one/proto"
)

const (
	port = ":50051"
)

// server is used to implement proto.GreeterServer.
type server struct {
	pb.UnimplementedServiceServer
}

func (s *server) UpdateFirstName(ctx context.Context, in *pb.UpdateFirstNameRequest) (*pb.UpdateFirstNameResponse, error) {
	response := &pb.UpdateFirstNameResponse{}

	_, err := database.DbConnection.Update("user_one").
		Set("first_name", in.FirstName).
		Where("id_users = ?", in.UsersId).
		Exec()

	if err != nil {
		response.Error = err.Error()
	}

	var id int
	_, err = database.DbConnection.Select("id").From("user_one").
		Where("id_users = ?", in.UsersId).
		Where("first_name = ?", in.FirstName).
		Load(&id)


	response.TableName = "user_one"
	response.RowId = int64(id)

	return response, nil
}

type Data struct {
	Column string `db:"COLUMN_NAME"`
}

func (s *server) Rollback(ctx context.Context, in *pb.RollbackRequest) (*pb.RollbackResponse, error){
	response := &pb.RollbackResponse{}
	var data []string


	_, err := database.DbConnection.Select("COLUMN_NAME").
		From(dbr.I("INFORMATION_SCHEMA.COLUMNS")).
		Where("TABLE_NAME = ?", in.TableName).
		Load(&data)

	if err != nil {
		return response, err
	}
	mapping := make(map[string]interface{})

	for i, _ := range data {
		if data[i] == "id" {
			data[i] = in.TableName+"_id"
		}
	}


	_, err = database.DbConnection.Select(data...).
		From(in.TableName+"_history").
		Where(fmt.Sprintf("%s_id = ?", in.TableName), in.Id).
		Limit(1).
		OrderDir("date_changed", false).
		Load(&mapping)


	mapping["id"] = mapping[in.TableName+"_id"]

	delete(mapping,in.TableName+"_id")

	builder := database.DbConnection.Update(in.TableName)
	for key, value := range mapping {
		if key != "id" {
			builder.Set(key, value)
		}
	}
	_, err = builder.Where("id = ?",mapping["id"]).Exec()
	if err != nil {
		return response, err
	}

	response.Success = true

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

