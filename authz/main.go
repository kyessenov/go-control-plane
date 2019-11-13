// Copyright 2018 Envoyproxy Authors
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

// Package main contains the test driver for testing xDS manually.
package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	authz "github.com/envoyproxy/go-control-plane/envoy/service/auth/v2"
)

type impl struct{}

func (_ *impl) Check(ctx context.Context, req *authz.CheckRequest) (*authz.CheckResponse, error) {
	log.Printf("received authz request: path: %v\n", req.Attributes.Request.Http.Path)
	return &authz.CheckResponse{}, nil
}

// main returns code 1 if any of the batches failed to pass all requests
func main() {
	grpcServer := grpc.NewServer()
	lis, err := net.Listen("tcp", "127.0.0.1:9091")
	if err != nil {
		log.Fatal(err)
	}

	authz.RegisterAuthorizationServer(grpcServer, &impl{})
	log.Println("access log server listening on 9091")

	if err = grpcServer.Serve(lis); err != nil {
		log.Println(err)
	}
}
