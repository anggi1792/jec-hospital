FROM golang:alpine as builder

WORKDIR /app

COPY . .

RUN go build -o app cmd/grpc/main.go

FROM scratch

WORKDIR /app

COPY --from=builder /app/app .

CMD [ "./app" ]

# bild image
# docker build -t <nama_image:tag>.

# runnning contaner
# docker run --name <container_name> -p <port:port> <name_image>

# running postgres container
# docker run --name pg-jec -p 5444:5432 -e POSTGRES_PASSWORD=jec-pass -e POSTGRES_USER=jec-user -e POSTGRES_DB=jec-db postgres
