#Base Image
FROM golang AS builder

# Working Directory
WORKDIR /app

#copy all fie
COPY . .

#running comment
RUN CGO_ENABLE=0 GOOS=linux go build -o main ./cmd/grpc/main.go

FROM scratch

WORKDIR /app

COPY --from=builder /app/main .



CMD ["./main"]

# bild image
# docker build -t <nama_image:tag>.

# runnning contaner
# docker run --name <container_name> -p <port:port> <name_image>

# running postgres container
# docker run --name pg-jec -p 5444:5432 -e POSTGRES_PASSWORD=jec-pass -e POSTGRES_USER=jec-user -e POSTGRES_DB=jec-db postgres
