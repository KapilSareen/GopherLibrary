FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy && go mod vendor

COPY . .

RUN go build -o main .

EXPOSE 8001

CMD ["./main"]
