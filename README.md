# go_microservices
Repository to build microservices with golang

set -o allexport; source ./.env ; set +o allexport

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pkg/protos/user/user.proto
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pkg/protos/dbmanager/dbmanager.proto