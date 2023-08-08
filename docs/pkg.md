<!-- Package List -->
# go get -u github.com/golang/protobuf/protoc-gen-go
# go get -u github.com/golang/protobuf@v1.5.0
# go get -u google.golang.org/grpc@v1.56.0

<!-- Command -->
<!-- ! Generate Proto -->
# protoc --go_out=. ./domain/hospital/proto/*.proto
<!-- ! Generate Proto GRPC -->
# protoc --go_out=. ./domain/hospital/proto/*.proto --go-grpc_out=.