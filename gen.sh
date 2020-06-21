#!/bin/bash

# protoc calculator/calcpb/calc.proto --go_out=plugins=grpc:.

# protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.

protoc blog/blogpb/blog.proto --go_out=plugins=grpc:.
