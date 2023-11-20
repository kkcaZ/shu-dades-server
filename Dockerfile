FROM golang:1.21-alpine3.18

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN mkdir -p /bin
RUN go build -v -o /bin ./cmd/main.go

ENV HOST=0.0.0.0
EXPOSE 8080

CMD ["/bin/main"]