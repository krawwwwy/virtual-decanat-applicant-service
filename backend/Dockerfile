FROM golang:1.23-alpine

WORKDIR /applicant

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

CMD ["./main"]

EXPOSE 8082