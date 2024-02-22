FROM golang:1.22-alpine as builder

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /go/app ./...

FROM alpine:latest

USER 1000

WORKDIR /go
COPY --from=builder /go/app .

CMD ["/go/app"]
