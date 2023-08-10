<!-- Package List -->
# go get -u github.com/golang/protobuf/protoc-gen-go
# go get -u github.com/golang/protobuf@v1.5.0
# go get -u google.golang.org/grpc@v1.56.0

<!-- Command -->
<!-- ! Generate Proto -->
# protoc --go_out=. ./domain/hospital/proto/*.proto
<!-- ! Generate Proto GRPC -->
# protoc --go_out=. ./domain/hospital/proto/*.proto --go-grpc_out=.


<!-- Step -->
# 0. hospital [Entity]
# 1. hospitalDto [Dto]
# 2. hospitalHandler

# 3. hospitalService
# 4. hospitalRepository [Connect to Database]


# docker build -t 172.16.80.157/hacktiv8-final-project/h_service_hospital:v0.0.1 -f deploy/docker/grpc/Dockerfile .
